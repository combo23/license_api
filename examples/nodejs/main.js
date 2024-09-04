const axios = require('axios');
const { machineIdSync } = require('node-machine-id');

const LICENSE_KEY = "TEST";

async function main() {
    try {
        console.log("Checking license...");

        const response = await axios.post(
            "http://127.0.0.1:8080/verify",
            {
                hwid: machineIdSync(),
                license_key: LICENSE_KEY
            },
            {
                headers: { "Content-Type": "application/json" }
            }
        );

        console.log(response.data.message);
        // verified 

    } catch (error) {
        if (error.response) {
            console.log(error.response.data.error);
            process.exit(1);
        } else {
            console.error(error);
            process.exit(1);
        }

    }
}

main();
