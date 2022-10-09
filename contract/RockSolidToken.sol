// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

contract RockSolidToken is ERC1155 {
    uint256 public constant Points = 1;
    uint256 public constant GoldBadge = 2;

    constructor() ERC1155("https://ipfs.io/ipfs/bafybeig6tvzn5thiqbspfz356vnma6v3xkzty6qevedp23wjiwu776h6wa/{id}.json") {
        _mint(msg.sender, Points, 1_000_000, "");
        _mint(msg.sender, GoldBadge, 500, "");
    }

    function uri(uint256 _tokenid) override public pure returns (string memory) {
        return string(
            abi.encodePacked(
                "https://ipfs.io/ipfs/bafybeig6tvzn5thiqbspfz356vnma6v3xkzty6qevedp23wjiwu776h6wa/",
                Strings.toString(_tokenid),".json"
            )
        );
    }
}
