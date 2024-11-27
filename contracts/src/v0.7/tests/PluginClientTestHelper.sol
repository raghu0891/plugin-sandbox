// SPDX-License-Identifier: MIT
pragma solidity ^0.7.0;

import "../PluginClient.sol";
import "../vendor/SafeMathPlugin.sol";

contract PluginClientTestHelper is PluginClient {
  using SafeMathPlugin for uint256;

  constructor(address _link, address _oracle) {
    setPluginToken(_link);
    setPluginOracle(_oracle);
  }

  event Request(bytes32 id, address callbackAddress, bytes4 callbackfunctionSelector, bytes data);
  event LinkAmount(uint256 amount);

  function publicNewRequest(
    bytes32 _id,
    address _address,
    bytes memory _fulfillmentSignature
  ) public {
    Plugin.Request memory req = buildPluginRequest(_id, _address, bytes4(keccak256(_fulfillmentSignature)));
    emit Request(req.id, req.callbackAddress, req.callbackFunctionId, req.buf.buf);
  }

  function publicRequest(
    bytes32 _id,
    address _address,
    bytes memory _fulfillmentSignature,
    uint256 _wei
  ) public {
    Plugin.Request memory req = buildPluginRequest(_id, _address, bytes4(keccak256(_fulfillmentSignature)));
    sendPluginRequest(req, _wei);
  }

  function publicRequestRunTo(
    address _oracle,
    bytes32 _id,
    address _address,
    bytes memory _fulfillmentSignature,
    uint256 _wei
  ) public {
    Plugin.Request memory run = buildPluginRequest(_id, _address, bytes4(keccak256(_fulfillmentSignature)));
    sendPluginRequestTo(_oracle, run, _wei);
  }

  function publicRequestOracleData(
    bytes32 _id,
    bytes memory _fulfillmentSignature,
    uint256 _wei
  ) public {
    Plugin.Request memory req = buildOperatorRequest(_id, bytes4(keccak256(_fulfillmentSignature)));
    sendOperatorRequest(req, _wei);
  }

  function publicRequestOracleDataFrom(
    address _oracle,
    bytes32 _id,
    address _address,
    bytes memory _fulfillmentSignature,
    uint256 _wei
  ) public {
    Plugin.Request memory run = buildOperatorRequest(_id, bytes4(keccak256(_fulfillmentSignature)));
    sendOperatorRequestTo(_oracle, run, _wei);
  }

  function publicCancelRequest(
    bytes32 _requestId,
    uint256 _payment,
    bytes4 _callbackFunctionId,
    uint256 _expiration
  ) public {
    cancelPluginRequest(_requestId, _payment, _callbackFunctionId, _expiration);
  }

  function publicPluginToken() public view returns (address) {
    return pluginTokenAddress();
  }

  function publicFulfillPluginRequest(bytes32 _requestId, bytes32) public {
    fulfillRequest(_requestId, bytes32(0));
  }

  function fulfillRequest(bytes32 _requestId, bytes32) public {
    validatePluginCallback(_requestId);
  }

  function publicPLI(uint256 _amount) public {
    emit LinkAmount(PLI_DIVISIBILITY.mul(_amount));
  }

  function publicOracleAddress() public view returns (address) {
    return pluginOracleAddress();
  }

  function publicAddExternalRequest(address _oracle, bytes32 _requestId) public {
    addPluginExternalRequest(_oracle, _requestId);
  }
}
