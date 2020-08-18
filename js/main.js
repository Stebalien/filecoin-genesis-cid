const CID = require('cids')
const multihash = require('multihashes')

const genesisMultihashString = '1220107d821c25dc0735200249df94a8bebc9c8e489744f86a4ca8919e81f19dcd72'
const genesisCID = new CID(1, 'dag-cbor', multihash.fromHexString(genesisMultihashString))

console.log("Genesis CID:", genesisCID.toString())
