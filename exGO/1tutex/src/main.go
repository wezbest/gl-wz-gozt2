/*
   .:        .       ..Oddd' ...o:':,':d:;,cl,';'....:'O. .d,c...:k......l.
    .;        .     cdd:;; . ..';.'.:cc:llc.;;odddc;.':dl  .;o..''lddolxx.
     .;       .. .oxl.....   ..,',;cc::;.......::d0kx:,'dd'.cd.';,olloxd.
       ''      . :0,..   . .:::;;;cc.....  ..     .;:;ccloxOk:,c';occol       .o
        .;      .x;..  ;:::;                              .ccol:o;::d;       cOo
 ;xdodddood.      :o;c:.                ..                    'oOOdc.       locd
 xk.     .lc,     ..;...cc'             ..             ':l''.''.:...      'xl;'.
  ;dddoddxdox,   .l0l::;;c.             ..             ,o::::ccdo,.....  c0xl,,:
             :. ;dl'';c:.              ....                .;;;',......  d .;ool
             : ..;ddc,oxk.           ...,''.            ;ldc.''';0l'... .KxdOKxk
             , ,0:,''odl             .',xx'..          .;lo;'.;dd'.,,....kXkc.
            ::c.;x:.'col            . .....                ;co,..,c0;.:c.: ....
            od..,cdk                 ...:;.. .            l0oc,,.ox:.;;d,;
          . ..    ,'               ....cO0:..   ...        ..  dd:.,'. o;o
         ...                  .... ....kKKo....               .dll,     ;,
         ...                     .. ...,00:..... .                        .. .
        ..                         .,'',cx:'..;.                           ...'d
       ..,                        ,:.....:  .,...                          ..  O
        ..                       ,;    .,.    ..,                              o
       .'                       ,:    ...      ...                          .  O
      ...                      ,,.   ...         .                         ;:c:0
     ..                       ..     . .           .            d::c:,c:clcloddX

Exercises from ztpanty all of it here
*/

package main

import (
	l "tut/lib"
)

var (
	banner ="ZTMPANTY All Lessons below"
)

func main() {
	// banner function
	l.ColoBanner(banner)

	l.Headr("1 - Name Prining")
	l.Fu1()

  l.Headr("2 - Type Declaration")
  l.Fu2()

  l.Headr("3 - Direct Assignment")
  l.Fu3()

  l.Headr("4 - Uninitialized Variable")
  l.Fu4()

  l.Headr("5 - Multiple Assignments")
  l.Fu5()

  l.Headr("6 - OK Error Idiom")
  l.Fu6()

  l.Headr("7 - Variable Reassignment")
  l.Fu7()

  l.Headr("8 - Block Assigment")
  l.Fu8()

  l.Headr("9 - Assignment and ignoring variables")
  l.Fu9()

}