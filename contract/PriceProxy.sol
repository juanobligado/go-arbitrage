/**
 * Proxy interface for reading Uniswap Compatible Pair prices
*/
pragma solidity >=0.5.0;
import  "./IUniswapV2Pair.sol";

contract UniswapView {
    function viewPair(address[] calldata _pair) external view returns(uint112[] memory){
        uint pairs_count = _pair.length;
        uint112[] memory out = new uint112[](2*pairs_count);
        for(uint i = 0;i < pairs_count ; i++){
            (uint112 t0_l,uint112 t1_l,) = IUniswapV2Pair(_pair[i]).getReserves();
            out[2*i] = t0_l;
            out[2*i + 1] = t1_l;
        }
        return out;
    }

}