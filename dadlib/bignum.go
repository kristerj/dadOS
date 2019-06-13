package dadlib



const bigzero = `
    000000000     
  00:::::::::00   
00:::::::::::::00 
0:::::::000:::::::0
0::::::0   0::::::0
0:::::0     0:::::0
0:::::0     0:::::0
0:::::0 000 0:::::0
0:::::0 000 0:::::0
0:::::0     0:::::0
0:::::0     0:::::0
0::::::0   0::::::0
0:::::::000:::::::0
 00:::::::::::::00 
   00:::::::::00   
     000000000     
`

const bigone = `
  1111111   
 1::::::1   
1:::::::1   
111:::::1   
   1::::1   
   1::::1   
   1::::1   
   1::::l   
   1::::l   
   1::::l   
   1::::l   
   1::::l   
111::::::111
1::::::::::1
1::::::::::1
111111111111
`

const bigtwo = `
222222222222222    
2:::::::::::::::22  
2::::::222222:::::2 
2222222     2:::::2 
            2:::::2 
            2:::::2 
         2222::::2  
    22222::::::22   
   22::::::::222     
2:::::22222        
2:::::2             
2:::::2             
2:::::2       222222
2::::::2222222:::::2
2::::::::::::::::::2
22222222222222222222
`
const bigthree = `
 333333333333333   
3:::::::::::::::33 
3::::::33333::::::3
3333333     3:::::3
            3:::::3
            3:::::3
    33333333:::::3 
    3:::::::::::3  
    33333333:::::3 
            3:::::3
            3:::::3
            3:::::3
3333333     3:::::3
3::::::33333::::::3
3:::::::::::::::33 
 333333333333333   
`
const bigfour = `
       444444444  
      4::::::::4  
     4:::::::::4  
    4::::44::::4  
   4::::4 4::::4  
  4::::4  4::::4  
 4::::4   4::::4  
4::::444444::::444
4::::::::::::::::4
4444444444:::::444
          4::::4  
          4::::4  
          4::::4  
        44::::::44
        4::::::::4
        4444444444
`
const bigfive = `
555555555555555555 
5::::::::::::::::5 
5::::::::::::::::5 
5:::::555555555555 
5:::::5            
5:::::5            
5:::::5555555555   
5:::::::::::::::5  
555555555555:::::5 
            5:::::5
            5:::::5
5555555     5:::::5
5::::::55555::::::5
 55:::::::::::::55 
    55:::::::::55   
      555555555     
`
const bigsix = `
        66666666   
       6::::::6    
      6::::::6     
     6::::::6      
    6::::::6       
   6::::::6        
  6::::::6         
 6::::::::66666    
6::::::::::::::66  
6::::::66666:::::6 
6:::::6     6:::::6
6:::::6     6:::::6
6::::::66666::::::6
 66:::::::::::::66 
   66:::::::::66   
     666666666     
`
const bigseven = `
77777777777777777777
7::::::::::::::::::7
7::::::::::::::::::7
777777777777:::::::7
           7::::::7 
          7::::::7  
         7::::::7   
        7::::::7    
       7::::::7     
      7::::::7      
     7::::::7       
    7::::::7        
   7::::::7         
  7::::::7          
 7::::::7           
77777777          
`
const bigeight = `
     888888888     
   88:::::::::88   
 88:::::::::::::88 
8::::::88888::::::8
8:::::8     8:::::8
8:::::8     8:::::8
 8:::::88888:::::8 
  8:::::::::::::8  
 8:::::88888:::::8 
8:::::8     8:::::8
8:::::8     8:::::8
8:::::8     8:::::8
8::::::88888::::::8
 88:::::::::::::88 
   88:::::::::88   
     888888888
`
const bignine = `
     999999999     
   99:::::::::99   
 99:::::::::::::99 
9::::::99999::::::9
9:::::9     9:::::9
9:::::9     9:::::9
 9:::::99999::::::9
  99::::::::::::::9
    99999::::::::9 
         9::::::9  
        9::::::9   
       9::::::9    
      9::::::9     
     9::::::9      
    9::::::9       
   99999999    
`
//Bignum accepts a number and returns a string of that number, but really big. 
//Its a good name for it.
func Bignum(num int) string {
	var bignums [10]string
	bignums[0] = bigzero
	bignums[1] = bigone
	bignums[2] = bigtwo
	bignums[3] = bigthree
	bignums[4] = bigfour
	bignums[5] = bigfive
	bignums[6] = bigsix
	bignums[7] = bigseven
	bignums[8] = bigeight
	bignums[9] = bignine

	return bignums[num]
}