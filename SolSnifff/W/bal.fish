function bal 

    # This function is for getting the balances of all of your accounts 
    # Wallets set 
    # w1 0x76021F7C89C43e7aEf0F58e5A72c582F0A8cf48B
    # w2 0x7434651d968DbC9CA32268Abbd63EC9Cfb447B31
   

    # Set Variables - Just Jeep changin this all the
    # Wallets 
    set A1 0x76021F7C89C43e7aEf0F58e5A72c582F0A8cf48B
    set A2 0x7434651d968DbC9CA32268Abbd63EC9Cfb447B31
    
    # Set RPC here 
    set GOE https://eth-goerli.g.alchemy.com/v2/wfowE284fYsqh-laeF5qmjOTQ3YTwuX_  
    set MUM https://polygon-mumbai.g.alchemy.com/v2/DlxcFxU5u-OioYGMvKhwWsAzAknhvQYd
    set SEP https://sepolia.infura.io/v3/4d9f7fa54ce44d1aa3319dca50aa3dd7
    set BSC https://sparkling-boldest-brook.bsc-testnet.discover.quiknode.pro/9c0a1ccf523e238a5e6d574a36ea192f5fcfb747

    # Actual Code starts here 

    clear
    # Displaying todays rates 
    # Today Eth Rate
    set today_eth_rate (curl -s rate.sx/1ETH)
    set today_bnb_rate (curl -s rate.sx/1BNB)
    
    # Aesthetic Display
    echo "---"
    echo "Todays ETH Rate" $today_eth_rate
    echo "Todays BNB Rate" $today_bnb_rate
    echo ""
    echo (set_color ff0000)"############################ Balances ##################################"
    echo ""
    echo (set_color 0000ff)"***********************************************************************"
    echo (set_color 0000ff)""
    echo (set_color 0000ff)"WRL1  : 0x76021F7C89C43e7aEf0F58e5A72c582F0A8cf48B "
    echo (set_color 0000ff)"WRL1K : 0x31a07a6705591eb67e83a018eee03e46447d81847149c70972c6eb348ec4190b"
    echo (set_color 0000ff)""
    echo (set_color 0000ff)"***********************************************************************"
    set w1g (cast balance $A1 --rpc-url $GOE ) 
    set w1m (cast balance $A1 --rpc-url $MUM ) 
    set w1s (cast balance $A1 --rpc-url $SEP ) 
    set w1b (cast balance $A1 --rpc-url $BSC ) 
    echo (set_color 0000ff)"Goerli    :" $w1g
    echo (set_color 0000ff)"Matic     :" $w1m
    echo (set_color 0000ff)"Sepolia   :" $w1s
    echo (set_color 0000ff)"Binance   :" $w1b
    echo ""
    echo (set_color ff0000)"-----------------------------------------------------------------------"
    echo " "
    echo (set_color ee82ee)"***********************************************************************"
    echo (set_color ee82ee)""
    echo (set_color ee82ee)"WRL2  : 0x7434651d968DbC9CA32268Abbd63EC9Cfb447B31"
    echo (set_color ee82ee)"WRL2K : 0xe97c3c7130b94bc6f64b98d8f1d195ab5e49dea477eed60934d217ce9cd58d27"
    echo (set_color ee82ee)""
    echo (set_color ee82ee)"***********************************************************************"
    set w1g (cast balance $A2 --rpc-url $GOE )
    set w1m (cast balance $A2 --rpc-url $MUM )
    set w1s (cast balance $A2 --rpc-url $SEP )
    set w1b (cast balance $A2 --rpc-url $BSC )
    echo (set_color ee82ee)"Goerli    :" $w1g
    echo (set_color ee82ee)"Matic     :" $w1m
    echo (set_color ee82ee)"Sepolia   :" $w1s
    echo (set_color ee82ee)"Binance   :" $w1b
    echo ""
    echo (set_color ff0000)"-------------------------------------------------------------------------"
    echo ""
    echo ""
    echo (set_color ff0000)"-------------------------------------------------------------------------"
    echo " "
    echo (set_color ff0000)"################################ END ####################################"


end