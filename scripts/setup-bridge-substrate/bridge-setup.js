const { WsProvider, ApiPromise } = require('@polkadot/api');
const { Keyring } = require("@polkadot/keyring");
const { BN } = require("bn.js");

const keyring = new Keyring({ type: 'sr25519' });

const wss_url = process.env.WSS_URL;
const seed = process.env.SUDO;
const relays = process.env.BRIDGE_RELAYS.split(',').map(s => {return s.trim()});
const resource_id = process.env.RESOURCE_ID;
const chain_seed = process.env.CHAIN_SEED;
const job_holdings_seed = process.env.JOB_HOLDINGS_SEED;

var sudo_changes = 0;

async function main() {
    const args = process.argv.slice(2);

    const provider = new WsProvider(wss_url);
    const api = await ApiPromise.create({provider});
    await api.isReady;
    const sudo_account = keyring.createFromUri(seed);
    const chain_account = keyring.createFromUri(chain_seed);
    const job_holdings_account = keyring.createFromUri(job_holdings_seed);
    const target_relayer_threshold = 0x01
    const whitelist_id = 5

    const [chain, nodeName, nodeVersion] = await Promise.all([
      api.rpc.system.chain(),
      api.rpc.system.name(),
      api.rpc.system.version()
    ]);

    console.log(`You are connected to chain ${chain} using ${nodeName} v${nodeVersion}`);
    console.log();

    var target_keys = {}
    for (const relay of relays) {
      target_keys[relay] = api.query.chainBridge.relayers.key(relay);
    }

    const current_keys = await api.query.chainBridge.relayers.keys()
      .then(r => {return r.map(r => {return r.toHex()})})

    for (const current_key of current_keys) {
      if (!Object.values(target_keys).includes(current_key)) {
        console.log(`Removing unknown relay key: ${current_key}`)
        var nonce = await api.rpc.system.accountNextIndex(sudo_account.address);
        await do_sudo(sudo_account, api, "remove_relayer", api.tx.chainBridge.removeRelayer(current_key), nonce, current_key);
      }
    }

    for (const relay in target_keys) {
      const target_key = target_keys[relay];
      if (!current_keys.includes(target_key)) {
        console.log(`Adding new relay: ${relay} [key=${target_key}]`)
        var nonce = await api.rpc.system.accountNextIndex(sudo_account.address);
        await do_sudo(sudo_account, api, "add_relayer", api.tx.chainBridge.addRelayer(relay), nonce, relay);
      } else {
        console.log(`Relay already added: ${relay}`)
      }
    }

    const resource_current_keys = await api.query.chainBridge.resources.keys()
      .then(r => {return r.map(r => {return r.toHex()})})
    
    resource_target = api.query.chainBridge.resources.key(resource_id);

    for (const current_resource_key of resource_current_keys) {
      if (current_resource_key != resource_target) {
        console.log(`Removing unknown resource key: ${current_resource_key}`)
        var nonce = await api.rpc.system.accountNextIndex(sudo_account.address);
        await do_sudo(sudo_account, api, "remove_resource", api.tx.chainBridge.removeResource(current_resource_key), nonce, current_resource_key);
      }
    }

    if (!resource_current_keys.includes(resource_target)) {
        console.log(`Setting resource target: ${resource_id} [Exchange.transfer]`)
        var nonce = await api.rpc.system.accountNextIndex(sudo_account.address);
        await do_sudo(sudo_account, api, "set_resource", api.tx.chainBridge.setResource(resource_id, "Exchange.transfer"), nonce, resource_id);
    } else {
        console.log(`Resource target already set ${resource_id} [Exchange.transfer]`)
    }

    const relayer_threshold = await api.query.chainBridge.relayerThreshold()
    if (relayer_threshold != target_relayer_threshold) {
        console.log(`Setting relayer threshold to: ${target_relayer_threshold}`);
        var nonce = await api.rpc.system.accountNextIndex(sudo_account.address);
        await do_sudo(sudo_account, api, "set_threshold", api.tx.chainBridge.setThreshold(target_relayer_threshold), nonce, target_relayer_threshold);
    } else {
        console.log(`Relayer threshold already set to ${target_relayer_threshold}`)
    }

    const current_chain_account = await api.query.chainBridge.chainAddress().then(addr => {return addr.toString()})
    if (chain_account.address != current_chain_account) {
        console.log(`Setting chain account to: ${chain_account.address}`)
        var nonce = await api.rpc.system.accountNextIndex(sudo_account.address);
        await do_sudo(sudo_account, api, "set_chainbridgeaccount", api.tx.chainBridge.setChainbridgeAccount(chain_account.address), nonce, chain_account.address);
    } else {
        console.log(`Chain account already set to ${chain_account.address}`);
    }

    const current_job_holdings_account = await api.query.jobs.jobHoldings().then(addr => {return addr.toString()})
    if (current_job_holdings_account != job_holdings_account.address) {
        console.log(`Setting jobholdings account to: ${job_holdings_account.address}`);
        var nonce = await api.rpc.system.accountNextIndex(sudo_account.address);
        await do_sudo(sudo_account, api, "set_holding_account", api.tx.jobs.setHoldingAccount(job_holdings_account.address), nonce, job_holdings_account.address);
    } else {
        console.log(`Job holdings account already set to ${job_holdings_account.address}`);
    }

    if (sudo_changes) {
        var nonce = await api.rpc.system.accountNextIndex(sudo_account.address);
        await do_sudo(sudo_account, api, "whitelist_chain", api.tx.chainBridge.whitelistChain(whitelist_id), nonce, whitelist_id);
    }

    console.log()

    const { tokenDecimals : token_decimals, tokenSymbol : token_symbol } = await api.registry.getChainProperties().toHuman()

    var bal = await api.query.system.account(chain_account.address);
    console.log(`chain account balance [${chain_account.address}]: ${bal.data.free.toLocaleString()/(10**token_decimals)} ${token_symbol}`);

    var bal = await api.query.system.account(sudo_account.address);
    console.log(`job holdings account balance [${sudo_account.address}]: ${bal.data.free.toLocaleString()/(10**token_decimals)} ${token_symbol}`);

    var bal = await api.query.system.account(job_holdings_account.address);
    console.log(`sudo account balance [${job_holdings_account.address}]: ${bal.data.free.toLocaleString()/(10**token_decimals)} ${token_symbol}`);

    for (const relay of relays) {
      const bal = await api.query.system.account(relay);
      console.log(`relay account balance [${relay}]: ${bal.data.free.toLocaleString()/(10**token_decimals)} ${token_symbol}`);
    }
}

function do_sudo(account, api, name, thing_to_do, nonce, id) {
    sudo_changes++;
    return new Promise((resolve, reject) => {
      api.tx.sudo.sudo(
          thing_to_do
      )
      .signAndSend(account, { nonce }, ({ status, events }) => {
          events.forEach(({ event: { data, method, section }, phase }) => {
            console.log('\t', phase.toString(), `: ${section}.${method}`, data.toString());
          });
          if (status.isInBlock) {
            console.log('Included at block hash', status.asInBlock.toHex());
          }
          if (status.isFinalized) {
            console.log('Finalized block hash', status.asFinalized.toHex());
            events
              .filter(({ event }) =>
                api.events.sudo.Sudid.is(event)
              )
              .forEach(({ event : { data: [result] } }) => {
                if (result.isError) {
                  let error = result.asError;
                  if (error.isModule) {
                    const decoded = api.registry.findMetaError(error.asModule);
                    const { docs, name, section } = decoded;

                    reject(`${section}.${name}: ${docs.join(' ')}`);
                  } else {
                    reject(error.toString());
                  }
                } else {
                    console.log("No error, we're done here....");
                }
              });
              resolve(id);
          }
    })
  });
}

function delay(ms) {
  console.log("delay", ms)
  return new Promise(r => setTimeout(r, ms))
}


main()
.catch((err) => {
  console.log("Error", err)
  process.exit(1)
})
.finally(() => {process.exit()})
