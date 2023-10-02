#!/bin/fish 
# Section for seting variables
set SEP https://sepolia.infura.io/v3/4d9f7fa54ce44d1aa3319dca50aa3dd7
set MUM https://polygon-mumbai.g.alchemy.com/v2/DlxcFxU5u-OioYGMvKhwWsAzAknhvQYd
set W1 0x76021F7C89C43e7aEf0F58e5A72c582F0A8cf48B
set W1K 0x31a07a6705591eb67e83a018eee03e46447d81847149c70972c6eb348ec4190b
set W2 0x7434651d968DbC9CA32268Abbd63EC9Cfb447B31
set W2K 0xe97c3c7130b94bc6f64b98d8f1d195ab5e49dea477eed60934d217ce9cd58d27

# This Section for setting ASCII Image 

set wallet1text " 
....                          ...',,,,'...;..':dM0l. ..,lWN:. . '      ox       
'.......                      ....,,''.;.:;....;XMk;'...ckNMk;'..      dx       
,'.... .    .....               ..',''..ll'  ..:k0Nl;...,;lkKWN0k      cd:;;;;;:
oc::,'...........              ..''';' ;xl  ...ckoKN:;,;,.,,;ckXK      ;;       
XWXOoc,'''...                 ...'.,c,,0d   ..'l0NNMdl:ll::''''xK      cc       
 .lOXNN0dc;',.....  ...         .''lcoNd   ..,xKKoXMKOkdo:c,,c:0K      ll       
     .:xKWWKkdc::;''.......  ....,:loX: . .dOKKd;.cMM0kddOXWWNNKx.     ;o       
          ,d0NWKOdl:;,''''..,''''::O0,,okOKX,oc,'.xKMxOOMMWc...'.      :o       
              .lOXNNNkxkdllc::;lk0XKKKk00;l:,'....K0Mx,c0XWNx.';       cd       
                 cdXMXOO00KXO0K0xoo:co:;;,.. ..'.'kcMk .:ldKMXl,.......;o..     
                .MNdOk;,.'Ol.....''...,::'',':ll,,dcNO   ':;oNN.       .,       
                 KWllo::,lK. .... .':odlc::;;...'.:XNN. .:cc.,NK.      .,       
                 dMX,c:,;'0: ..';dO00KKKK0OOOoxklc.;NM0xxkKdl.xM.      .'       
  ..             .WNX0O,,,d0kO0KKOxll:,...;cdk0NXXO;0MWox;lOl,kMo      ..       
                  0WXNNXXNXOxoccodkOk'. ......':dOWWWMNlkX;;xoOKW.     .lllooooo
                 .cMWNNXOol:;,,,,,clk0O'.    ..'':lKNWX,kNN,.;ldM       .       
                 .dMW0k:,;'.......'.,lxXO,     ..,,dKWl.dkWW;..dMddddddxc       
               .oNM0l;'.....    ....',;lkKo     .''cKM'.,dlNX. cM:::;;:kK       
.             cNWk:........         .,,;;d00,. ....lKM.ldxolKl.,M      dK       
...         ,KMXl,.........          ..,.'l:Kd.....ONN,l.:d;xk.,MkOOOOOOO       
.         .OMXx,'.....              . .'..'c.ll..,lOMk o..;dOd.xM      .d       
        .xWXo;.......            .   ..... ,'.;:':OWK; l'..'OXNKdcccccclc       
       oWNx:.....   .             .   ''....,..,;xWN'o,:c..''XW,                
     'XWO:..... .   .                 .;.. .''.'cWN.  ,.c '. ,XW.              

 ██╗    ██╗  █████╗  ██╗      ██╗      ███████╗ ████████╗  ██╗    
 ██║    ██║ ██╔══██╗ ██║      ██║      ██╔════╝ ╚══██╔══╝ ███║    
 ██║ █╗ ██║ ███████║ ██║      ██║      █████╗      ██║    ╚██║    
 ██║███╗██║ ██╔══██║ ██║      ██║      ██╔══╝      ██║     ██║    
 ╚███╔███╔╝ ██║  ██║ ███████╗ ███████╗ ███████╗    ██║     ██║    
  ╚══╝╚══╝  ╚═╝  ╚═╝ ╚══════╝ ╚══════╝ ╚══════╝    ╚═╝     ╚═╝    

███████╗ ███╗   ██╗ ██╗ ███████╗ ███████╗     ██████╗   █████╗  ███╗   ██╗ ████████╗ ██╗   ██╗
██╔════╝ ████╗  ██║ ██║ ██╔════╝ ██╔════╝     ██╔══██╗ ██╔══██╗ ████╗  ██║ ╚══██╔══╝ ╚██╗ ██╔╝
███████╗ ██╔██╗ ██║ ██║ █████╗   █████╗       ██████╔╝ ███████║ ██╔██╗ ██║    ██║     ╚████╔╝ 
╚════██║ ██║╚██╗██║ ██║ ██╔══╝   ██╔══╝       ██╔═══╝  ██╔══██║ ██║╚██╗██║    ██║      ╚██╔╝  
███████║ ██║ ╚████║ ██║ ██║      ██║          ██║      ██║  ██║ ██║ ╚████║    ██║       ██║   
╚══════╝ ╚═╝  ╚═══╝ ╚═╝ ╚═╝      ╚═╝          ╚═╝      ╚═╝  ╚═╝ ╚═╝  ╚═══╝    ╚═╝       ╚═╝   
"

set wallet2text " 
W..,klo:c:cK0                                                ':. .      ...  .'M
Mx.'l0dd'c;kK.                                            .  ,c.   . .        .M
XlKXOldc.oN0O. ..  ..                                     ...c:.  .           .M
Mx0Kx:kOOOo.........                          ...       .',;;l,,;cldxkkOOOOO000W
NX0K0x;.                                                             ....''',,'K
k,                                                     ..;,';;:cc::::;;;;,,'',,M
O,;:cldxxdl:,.       .           .:dOKXKKKKKXXXXXKNWNNWXdddlc;'...............,M
X........':lodxkkxxxkkO0k0Ok000KKX0kdc',;,',,,;:::dkXNd'''....................,M
N.. ..   .       ....   .:lkkN0xl, ..',;;:::,;,'.,ONd.......... . ............'M
W .                       ...:KNko'..'',,''','',:NO..         ... .           'M
W..  .                       .'xWX0doldxx0kkkddkWd          ... .     .       .M
W     .                         oMXx;;,,,;:clxKN:                             .M
M   ..                          .xWN;....'..'kk.                  ...         'M
M  .       ...                  ..X0O:.....,Kk..              . .      ...    'M
M                               ..,Xx:,  .lXXc..            ... .             'M
M                ....           . 'd0od.'oKOO..                ..  ...        'M
M                               . ..;dOldl.c'..               ..    .';,....  .M
M              ...                ...:do,  ....   .            .          ..  .M
M..         ....                  ...;0:. .....   .          .......          'M
M       ..                        ...oNll'.''.............     ......         .M
M           .         .       .. ..'cN;..l::llc::;,,;'......   .  .''.  .     .M
M                    .  ..    ....'oO:...';...:lldkkOOx;,,;;'....   .....  ....M
M                  ..   ..    ',;lKx.              ..,lkO00kol;,''.......  ...'M
M.    .          ...... ..  .'dkOx;...........            'lk0KOo::;,'........'M
██╗    ██╗  █████╗  ██╗      ██╗      ███████╗ ████████╗     ██████╗     
██║    ██║ ██╔══██╗ ██║      ██║      ██╔════╝ ╚══██╔══╝     ╚════██╗    
██║ █╗ ██║ ███████║ ██║      ██║      █████╗      ██║         █████╔╝    
██║███╗██║ ██╔══██║ ██║      ██║      ██╔══╝      ██║        ██╔═══╝     
╚███╔███╔╝ ██║  ██║ ███████╗ ███████╗ ███████╗    ██║        ███████╗    
 ╚══╝╚══╝  ╚═╝  ╚═╝ ╚══════╝ ╚══════╝ ╚══════╝    ╚═╝        ╚══════╝    

██████╗  ██████╗  ██╗ ███╗   ██╗ ██╗  ██╗     ██████╗  ██╗ ███████╗ ███████╗
██╔══██╗ ██╔══██╗ ██║ ████╗  ██║ ██║ ██╔╝     ██╔══██╗ ██║ ██╔════╝ ██╔════╝
██║  ██║ ██████╔╝ ██║ ██╔██╗ ██║ █████╔╝      ██████╔╝ ██║ ███████╗ ███████╗
██║  ██║ ██╔══██╗ ██║ ██║╚██╗██║ ██╔═██╗      ██╔═══╝  ██║ ╚════██║ ╚════██║
██████╔╝ ██║  ██║ ██║ ██║ ╚████║ ██║  ██╗     ██║      ██║ ███████║ ███████║
╚═════╝  ╚═╝  ╚═╝ ╚═╝ ╚═╝  ╚═══╝ ╚═╝  ╚═╝     ╚═╝      ╚═╝ ╚══════╝ ╚══════╝
"

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
# Sending to Sepolia 
# Section for sending the actual commands 
echo (set_color 97FEED)""
echo " ███████╗ ███████╗ ██████╗   ██████╗  ██╗      ██╗  █████╗ "
echo " ██╔════╝ ██╔════╝ ██╔══██╗ ██╔═══██╗ ██║      ██║ ██╔══██╗"
echo " ███████╗ █████╗   ██████╔╝ ██║   ██║ ██║      ██║ ███████║"
echo " ╚════██║ ██╔══╝   ██╔═══╝  ██║   ██║ ██║      ██║ ██╔══██║"
echo " ███████║ ███████╗ ██║      ╚██████╔╝ ███████╗ ██║ ██║  ██║"
echo " ╚══════╝ ╚══════╝ ╚═╝       ╚═════╝  ╚══════╝ ╚═╝ ╚═╝  ╚═╝"
echo ""
echo "Sending Mesages from W1-->W1 Self"
echo "--------------------------------------"
cast send \
    $W1 \
    --private-key $W1K \
    --rpc-url $SEP \
    $(cast --from-utf8 $wallet1text)
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
    $(cast --from-utf8 $wallet2text)
echo "DONE W1-->W1"
echo "--------------------------------------"
echo "--------------------------------------"
echo "--------------------------------------"

# Section for sending the actual commands 
echo (set_color FBFFB1)""
echo " ███╗   ███╗ ██╗   ██╗ ███╗   ███╗ ██████╗   █████╗  ██╗"
echo " ████╗ ████║ ██║   ██║ ████╗ ████║ ██╔══██╗ ██╔══██╗ ██║"
echo " ██╔████╔██║ ██║   ██║ ██╔████╔██║ ██████╔╝ ███████║ ██║"
echo " ██║╚██╔╝██║ ██║   ██║ ██║╚██╔╝██║ ██╔══██╗ ██╔══██║ ██║"
echo " ██║ ╚═╝ ██║ ╚██████╔╝ ██║ ╚═╝ ██║ ██████╔╝ ██║  ██║ ██║"
echo " ╚═╝     ╚═╝  ╚═════╝  ╚═╝     ╚═╝ ╚═════╝  ╚═╝  ╚═╝ ╚═╝"
echo ""
echo "Sending Mesages from W1-->W1 Self"
echo "--------------------------------------"
cast send \
    $W1 \
    --private-key $W1K \
    --rpc-url $MUM \
    $(cast --from-utf8 $wallet1text)
echo "DONE W1-->W1"
echo "--------------------------------------"
echo "--------------------------------------"
echo "--------------------------------------"
echo (set_color FF55BB)"Sending Mesages from W2-->W2 Self"
echo "--------------------------------------"
cast send \
    $W2 \
    --private-key $W2K \
    --rpc-url $MUM \
    $(cast --from-utf8 $wallet2text)
echo "DONE W1-->W1"
echo "--------------------------------------"
echo "--------------------------------------"
echo "--------------------------------------"

