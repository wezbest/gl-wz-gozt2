#!/bin/fish 
# Section for seting variables
set SEP https://sepolia.infura.io/v3/4d9f7fa54ce44d1aa3319dca50aa3dd7
set W1 0x76021F7C89C43e7aEf0F58e5A72c582F0A8cf48B
set W1K 0x31a07a6705591eb67e83a018eee03e46447d81847149c70972c6eb348ec4190b
set W2 0x7434651d968DbC9CA32268Abbd63EC9Cfb447B31
set W2K 0xe97c3c7130b94bc6f64b98d8f1d195ab5e49dea477eed60934d217ce9cd58d27

# Section for commands
clear
echo (set_color green)"Commands for sending arbitray messages"
echo "--------------------------------------"
echo "Command Structure"
echo "--------------------------------------"
echo "cast send "
echo "    <DestinationWallet> "
echo "    --private-key <OriginPk>"
echo "    --rpc-url <rpc>"

# Section for sending the actual commands 
echo "Sending Mesages from W1-->W1 Self"
echo "--------------------------------------"
cast send \
    $W1 \
    --private-key $W1K \
    --rpc-url $SEP \
    $(cast --from-utf8 "Wallet1 - SmellHerFarts")
echo "DONE W1-->W1"
echo "--------------------------------------"
echo "--------------------------------------"
echo "--------------------------------------"
echo "Sending Mesages from W2-->W2 Self"
echo "--------------------------------------"
cast send \
    $W2 \
    --private-key $W2K \
    --rpc-url $SEP \
    $(cast --from-utf8 "Wallet2 - DrinkherPiss")
echo "DONE W1-->W1"
echo "--------------------------------------"
echo "--------------------------------------"
echo "--------------------------------------"

