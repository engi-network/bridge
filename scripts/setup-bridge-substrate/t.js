const { WsProvider, ApiPromise } = require('@polkadot/api');
const { Keyring } = require("@polkadot/keyring");
const { BN } = require("bn.js");

var alice;

(async () => {
    const provider = new WsProvider("ws://127.0.0.1:9944");
    const api = await ApiPromise.create({provider});

    await api.isReady;

    // NOTE: will need a 'production' bridge account for substrate
    const keyring = new Keyring({ type: 'sr25519' });
    alice = keyring.createFromUri('//Alice');


    await do_sudo(api, "whitelist_chain", api.tx.chainBridge.whitelistChain(1));
    await do_sudo(api, "add_relayer", api.tx.chainBridge.addRelayer("5CVNdoFdQcc1WggXkfEHdzEMUrnKR7jYToY3ieuoVFWAmS4C"));
    await do_sudo(api, "set_resource", api.tx.chainBridge.setResource("0x00000000000000000000000000000063822bbd62abfb4ab9c92210c193e71b01", "Exchange.transfer"));
    await do_sudo(api, "set_threshold", api.tx.chainBridge.setThreshold(1));

    // Fund the relayer accounts
    const nonce = await api.rpc.system.accountNextIndex(alice.address);
    const unsub = await api.tx.balances.transfer("5CVNdoFdQcc1WggXkfEHdzEMUrnKR7jYToY3ieuoVFWAmS4C", 10000000000000)
        .signAndSend(alice, { nonce }, ({ status, events }) => {
            if (status.isInBlock || status.isFinalized) {
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

                      console.log(`${section}.${name}: ${docs.join(' ')}`);
                    } else {
                      console.log(error.toString());
                    }
                  } else {
                      console.log("No error, we're done here....");
                  }
                });
                console.log(`UNSUBBING FROM txfer`);
              unsub();
            }
        });

    // EXAMPLE Substrate to ETH transaction
    //await cb_sell(api, "2000000000000000000");
})()


async function cb_sell(api, amount) {
    const nonce = await api.rpc.system.accountNextIndex(alice.address);
    const unsub = await api.tx.exchange.sell("0xe9623f3ca3CcC1c1D415F3196D4B75007B316aC3", amount, 1)
        .signAndSend(alice, { nonce }, ({ status, events }) => {
            if (status.isInBlock || status.isFinalized) {
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

                      console.log(`${section}.${name}: ${docs.join(' ')}`);
                    } else {
                      console.log(error.toString());
                    }
                  } else {
                      console.log("No error, we're done here....");
                  }
                });
                console.log(`UNSUBBING FROM txfer`);
              unsub();
            }
        });
}

async function do_sudo(api, name, thing_to_do) {
    const nonce = await api.rpc.system.accountNextIndex(alice.address);
    const unsub = await api.tx.sudo.sudo(
        thing_to_do
    )
    .signAndSend(alice, { nonce }, ({ status, events }) => {
        if (status.isInBlock || status.isFinalized) {
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

                  console.log(`${section}.${name}: ${docs.join(' ')}`);
                } else {
                  console.log(error.toString());
                }
              } else {
                  console.log("No error, we're done here....");
              }
            });
            console.log(`UNSUBBING FROM ${name}`);
          unsub();
        }
    });
}
