#!/bin/fish 
echo ""
echo "****************************************************"
echo "This fish shell command will be sending transactions from "
echo "one wallet to another using cast"
echo "****************************************************"
echo "" 
echo (set_color blue)" Wallets & Keys"
echo (set_color blue)" ---------------" 
echo (set_color blue)"W1 = 0x76021F7C89C43e7aEf0F58e5A72c582F0A8cf48B"
echo (set_color blue)"W1K = 0x31a07a6705591eb67e83a018eee03e46447d81847149c70972c6eb348ec4190b"
echo (set_color blue)"W2 = 0x7434651d968DbC9CA32268Abbd63EC9Cfb447B31"
echo (set_color blue)"W2K = 0xe97c3c7130b94bc6f64b98d8f1d195ab5e49dea477eed60934d217ce9cd58d27"
echo (set_color blue)"---------------"
echo ""
echo ""
echo (set_color magenta)"Command Structure "
echo (set_color magenta)"//////////////////////////"
echo (set_color magenta)"cast send '\'"
echo (set_color magenta)"  --chain goerli '\'"
echo (set_color magenta)" --rpc-url $SEPOLIA_KEY '\'"
echo (set_color magenta)"  --private-key $GOERLI_P_KEY '\'"
echo (set_color magenta)"  $TW1 '\'"
echo (set_color magenta)"  --value 0.02ether"
echo (set_color magenta)"//////////////////////////"
echo ""
echo "" 
set SEP https://sepolia.infura.io/v3/4d9f7fa54ce44d1aa3319dca50aa3dd7
set W1k = 0x31a07a6705591eb67e83a018eee03e46447d81847149c70972c6eb348ec4190b
set W2K = 0xe97c3c7130b94bc6f64b98d8f1d195ab5e49dea477eed60934d217ce9cd58d27
echo "Sending 1 Ether from W1 --> W2"
echo "++++++++++++++++++++++++++++++++++++++++++" 
cast send \
    --chain sepolia \
    --rpc-url $SEP \
    --private-key $W1K \
    $W2 \
    --value 1.00ether
echo "DONE W1 ---> W2"
echo "++++++++++++++++++++++++++++++++++++++++++" 
echo ""
echo "Sending 1 Ether W1 <-- W2"
echo "++++++++++++++++++++++++++++++++++++++++++" 
cast send \
    --chain sepolia \
    --rpc-url $SEP \
    --private-key $W2K \
    $W1 \
    --value 1.00ether
echo "DONE W1 <--- W2"
echo "++++++++++++++++++++++++++++++++++++++++++" 
echo ""
