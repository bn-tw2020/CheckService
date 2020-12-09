const { FileSystemWallet, Gateway } = require('fabric-network');
const fs = require('fs');
const { resolve } = require('path');
const path = require('path');
const ccpPath = path.resolve(__dirname, '..', 'network' ,'connection.json');
const ccpJSON = fs.readFileSync(ccpPath, 'utf8');
const ccp = JSON.parse(ccpJSON);

async function send(type, func, args,res){ // type 1이면 invoke, type 0이면 query
    try{
        const walletPath = path.join(process.cwd(),'wallet')
        const wallet = new FileSystemWallet(walletPath);
        console.log(`wallet path: ${walletPath}`);

        const userExists = await wallet.exists('user1');
        if (!userExists) {
            console.log('An identity for the user "user1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'user1', discovery: { enabled: false } });

        const network = await gateway.getNetwork('mychannel');
        const contract = network.getContract('bloc');

        // type 1이면 invoke, type 0이면 query
        if(type) {
            await contract.submitTransaction(func, ...args);
            console.log('Transaction has been submitted');
            await gateway.disconnect();
            res.status(200).send('success');
        } else {
            const result = await contract.evaluateTransaction(func, ...args);
            console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
            // res.status(200).send(result.toString());
            let obj = JSON.parse(result);
            res.status(200).json(obj);
        }
    }catch(e){
        console.error(`Failed to submit transaction: ${e}`);
        res.send(`Failed to submit transaction: ${e}`);
    }
}

module.exports = {
    send
}