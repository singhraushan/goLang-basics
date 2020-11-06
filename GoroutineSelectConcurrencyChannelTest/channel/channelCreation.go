package main
import (
	"fmt"
"time"
)

func main(){
	var channel1 chan int;
	fmt.Println("default value of channel1: ",channel1)
	fmt.Printf("type of channel1:%T\n ",channel1)

	channel2 := make(chan int)
	fmt.Println("default value of channel2: ",channel2)
	fmt.Printf("type of channel2:%T\n ",channel2)

	fmt.Println("=======================")
	channel3 := make(chan int)
	go receiveChannelData(channel3)
	channel3 <- 25 //receive will wait for send
	fmt.Println("=======================")
	channel4 := make(chan string)
	go sendDataToCannel(channel4)
	for{
		 element,ok:= <- channel4 
		 if(ok==false){
			fmt.Println("channel4 is closed")
			break
		 }
		 fmt.Println("channel4 data:", element)
	}
	fmt.Println("=======================")
	sendReceiveDataToCannel()
	fmt.Println("=======================")

	channelLenCap()
	fmt.Println("=======================")
}

func receiveChannelData(ch1 chan int){
	fmt.Println(275+ <-ch1)// will wait untill ch1 get some data
}

func sendDataToCannel(mychnl chan string){
	for v := 0; v < 4; v++ { 
        mychnl <- "GeeksforGeeks"
    } 
    close(mychnl) 
}

func sendReceiveDataToCannel(){

	ch := make(chan string , 4)// passing capacity

   go func(){
	ch <- "GFG1"
	ch <- "GFG2"
	ch <- "GFG3"
	ch <- "GFG4"
	close(ch)
   }()

   time.Sleep(1*time.Second)
   fmt.Println("Length of the ch is: ", len(ch)) // if no time.Sleep(1*time.Second) then len might be 0 because might execute before 
   for val := range ch { 
	fmt.Println("val: ",val)
   }

}

func channelLenCap() { 
  
    // Creating a channel 
    // Using make() function 
    mychnl := make(chan string, 4) 
    mychnl <- "GFG"
    mychnl <- "gfg"
    mychnl <- "Geeks"
    mychnl <- "GeeksforGeeks"
  
    // Finding the length of the channel 
    // Using len() function 
	fmt.Println("Length of the channel is: ", len(mychnl)) 
	fmt.Println("Caoacity of the channel is: ", cap(mychnl)) 
} 

