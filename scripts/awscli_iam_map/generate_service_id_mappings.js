const fs = require("fs");
const all = require("aws-sdk");

require('aws-sdk/lib/maintenance_mode_message').suppress = true;

function isUppercase(word) {
    return /^\p{Lu}/u.test(word);
}
var sdkNames = Object.keys(all).filter((name) => {
    try {
        all[name].constructor.asdf
        all[name].serviceIdentifier.asdf
        // svc = new all[name]();
        // svc.api.serviceId
        return true
    } catch {
        return false
    }
})
console.log(sdkNames)

let value = Object.fromEntries(
    sdkNames.map((name) => {
        k = Object.keys(all.apiLoader.services[all[name].serviceIdentifier])[0]

        let versions = all.apiLoader.services[all[name].serviceIdentifier]
        let ver_name = Object.keys(versions)[0]
        let config = versions[ver_name]

        return [config.metadata.serviceId, name]
    })
);

fs.writeFileSync("svc_id_to_js_names.json", JSON.stringify(value, null, 2));