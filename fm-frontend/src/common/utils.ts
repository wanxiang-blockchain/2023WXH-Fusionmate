import { ethers } from "ethers";


function abiRawEncode(encTypes: any, encValues: any) {
  const hexStr = ethers.utils.defaultAbiCoder.encode(encTypes, encValues);
  return Buffer.from(hexStr.slice(2, hexStr.length), 'hex');
}

function keccak256(arg: any) {
  const hexStr = ethers.utils.keccak256(arg);
  return Buffer.from(hexStr.slice(2, hexStr.length), 'hex');
}

// Recursively finds all the dependencies of a type
function dependencies(primaryType: any, found: any = [], types: any = {}) {
  if (found.includes(primaryType)) {
    return found;
  }
  if (types[primaryType] === undefined) {
    return found;
  }
  found.push(primaryType);
  for (let field of types[primaryType]) {
    for (let dep of dependencies(field.type, found)) {
      if (!found.includes(dep)) {
        found.push(dep);
      }
    }
  }
  return found;
}

function encodeType(primaryType: any, types: any = {}) {
  // Get dependencies primary first, then alphabetical
  let deps = dependencies(primaryType);
  deps = deps.filter((t: any) => t != primaryType);
  deps = [primaryType].concat(deps.sort());

  // Format as a string with fields
  let result = '';
  for (let type of deps) {
    if (!types[type])
      throw new Error(`Type '${type}' not defined in types (${JSON.stringify(types)})`);
    result += `${type}(${types[type].map(({ name, type }: any) => `${type} ${name}`).join(',')})`;
  }
  return result;
}

function typeHash(primaryType: any, types: any = {}) {
  return keccak256(Buffer.from(encodeType(primaryType, types)));
}

function encodeData(primaryType: any, data: any, types: any = {}) {
  let encTypes = [];
  let encValues = [];

  // Add typehash
  encTypes.push('bytes32');
  encValues.push(typeHash(primaryType, types));

  // Add field contents
  for (let field of types[primaryType]) {
    let value = data[field.name];
    if (field.type == 'string' || field.type == 'bytes') {
      encTypes.push('bytes32');
      value = keccak256(Buffer.from(value));
      encValues.push(value);
    } else if (types[field.type] !== undefined) {
      encTypes.push('bytes32');
      value = keccak256(encodeData(field.type, value, types));
      encValues.push(value);
    } else if (field.type.lastIndexOf(']') === field.type.length - 1) {
      throw 'TODO: Arrays currently unimplemented in encodeData';
    } else {
      encTypes.push(field.type);
      encValues.push(value);
    }
  }

  return abiRawEncode(encTypes, encValues);
}

function domainSeparator(domain: any) {
  const types = {
    EIP712Domain: [
      { name: 'name', type: 'string' },
      { name: 'version', type: 'string' },
      { name: 'chainId', type: 'uint256' },
      { name: 'verifyingContract', type: 'address' },
      { name: 'salt', type: 'bytes32' }
    ].filter(a => domain[a.name])
  };
  return keccak256(encodeData('EIP712Domain', domain, types));
}

function structHash(primaryType: any, data: any, types = {}) {
  return keccak256(encodeData(primaryType, data, types));
}

function digestToSign(domain: any, primaryType: any, message: any, types: any = {}) {
  return keccak256(
    Buffer.concat([
      Buffer.from('1901', 'hex'),
      domainSeparator(domain),
      structHash(primaryType, message, types),
    ])
  );
}

export async function sign(domain: any, primaryType: any, message: any, types: any = {}, signer: any) {
  let signature;

  try {
    if (signer._signingKey) {
      const digest = digestToSign(domain, primaryType, message, types);
      signature = signer._signingKey().signDigest(digest);
      signature.v = '0x' + (signature.v).toString(16);
    } else {
      const address = await signer.getAddress();
      const msgParams = JSON.stringify({ domain, primaryType, message, types });

      // signature = await signer.signTypedData(domain, types, message)
      signature = await signer.provider.jsonRpcFetchFunc(
        'eth_signTypedData_v4',
        [address, msgParams]
      );
      // console.log(signature)
    }
  } catch (e) {
    throw e;
  }

  return signature;
}
