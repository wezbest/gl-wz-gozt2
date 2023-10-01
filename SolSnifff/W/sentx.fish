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
echo (set_color magenta)"  --value 0.1ether '\'"
echo (set_color magenta)"  <DestinationWallet> '\'"
echo (set_color magenta)"  --rpc-url <rpc> '\'"
echo (set_color magenta)"  <PkOfSender> '\'"
echo (set_color magenta)"  --value 0.02ether"
echo (set_color magenta)"//////////////////////////"
echo ""
echo "" 
set SEP https://sepolia.infura.io/v3/4d9f7fa54ce44d1aa3319dca50aa3dd7
set W1 0x76021F7C89C43e7aEf0F58e5A72c582F0A8cf48B
set W1K 0x31a07a6705591eb67e83a018eee03e46447d81847149c70972c6eb348ec4190b
set W2 0x7434651d968DbC9CA32268Abbd63EC9Cfb447B31
set W2K 0xe97c3c7130b94bc6f64b98d8f1d195ab5e49dea477eed60934d217ce9cd58d27
echo (set_color 97FEED)"Sending 1 Ether from W1 --> W2"
echo "++++++++++++++++++++++++++++++++++++++++++"
echo "++++++++++++++++++++++++++++++++++++++++++"
echo "++++++++++++++++++++++++++++++++++++++++++"
cast send \
    --value 0.1ether \
    $W2 \
    --rpc-url $SEP \
    --private-key $W1K
echo "DONE W1 ---> W2"
echo "Current Balance - Wallet 1"
cast balance $W1 --rpc-url $SEP
echo "++++++++++++++++++++++++++++++++++++++++++" 
echo "++++++++++++++++++++++++++++++++++++++++++" 
echo "++++++++++++++++++++++++++++++++++++++++++" 
echo ""
echo (set_color FBFFB1)"Sending 1 Ether W1 <-- W2"
echo "++++++++++++++++++++++++++++++++++++++++++" 
echo "++++++++++++++++++++++++++++++++++++++++++" 
cast send \
    --value 0.1ether \
    $W1 \
    --rpc-url $SEP \
    --private-key $W2K
echo "DONE W1 <--- W2"
echo "Current Balance - Wallet 2"
cast balance $W2 --rpc-url $SEP
echo "++++++++++++++++++++++++++++++++++++++++++" 
echo "++++++++++++++++++++++++++++++++++++++++++" 
echo "++++++++++++++++++++++++++++++++++++++++++" 
echo ""
