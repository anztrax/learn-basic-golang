package main

import "fmt"
import "math"
import "math/rand"
import "math/cmplx"
import "runtime"
import "time"
import "strings"
import "io";
import "image";

//this is function in go
func add(x int,y int) int{
	return x + y;
}

func add2(x, y int)int{
	return x + y;
}

//multiple result
func swap(x , y string)(string,string){
	return y,x;
}

//naked return and you can named return values (this is bad for readability
func split(sum int)(x, y int){
	x = sum * 10;
	y = sum / 10;
	return;
}

func testDeclaringVariable(){
	var c, python ,java bool;
	var javascript , typeScript = true, "Hmm Maybe !";
	const formatText = "%T(%v)\n"; //type : value

	c = true;
	python = false;
	java = true;

	var (
		toBe bool = false
		maxint uint64 = 1 << 64 -1
		z complex128 = cmplx.Sqrt(-5 + 121)
	)

	//zero values based
	var var1 int;
	var var2 float64;
	var var3 bool;
	var var4 string;
	var var5 int = -100;
	var var6 *int = nil;

	//translated into world :D
	const thisIsConstValue = "世界";

	fmt.Println("=============")
	fmt.Println("c, python, java, javascript, typeScript = ",c,python,java,javascript,typeScript);
	fmt.Printf(formatText,toBe,toBe);
	fmt.Printf(formatText,maxint,maxint);
	fmt.Printf(formatText,z,z);
	fmt.Println("zero based value (var1,var2,var3,var4,var6) : ",var1,var2,var3,var4,var6);
	fmt.Println("convert value to uint",uint(var5));

	fmt.Println("This is const value : ","hello",thisIsConstValue);
	fmt.Println("=============")
}

func testFlowControl(){
	fmt.Println("===============");

	//for
	sumVal := 0;
	for i:=0; i < 10 ;i++{
		sumVal += i;
	}
	fmt.Println("sum value : ",sumVal);

	//lol while
	sumVal = 1;
	for sumVal < 100 {
		sumVal += sumVal;
		fmt.Println(sumVal);
	}

	fmt.Println(sqrt(4),sqrt(-4));
	fmt.Println("pow with max: ");
	fmt.Println("===============");
	fmt.Println(
		powWithMax(3,3,20),
		powWithMax(2,2,100),	//if multiline using , at the end of line
	);

	fmt.Println("Go runs on");
	switch os := runtime.GOOS; os{
	case "darwin":
		fmt.Println("OSX");
	case "linux":
		fmt.Println("Linux");
	default:
		fmt.Printf("%s",os);
	}

	fmt.Println("When's Saturday ?");
	today := time.Now().Weekday();
	switch time.Saturday {
		case today + 0:
			fmt.Println("Today.");
		case today + 1:
			fmt.Println("Tomorrow.");
		case today + 2:
			fmt.Println("In 2 days.");
		default:
		fmt.Println("Too Far Away.");
	}

	//if write long if -else you can use switch (true) ?
	todayTime := time.Now();
	switch {		//this is switch true
	case todayTime.Hour() < 12:
		fmt.Println("Good Morning !");
	case todayTime.Hour() < 17:
		fmt.Println("Good Afternoon !");
	default:
		fmt.Println("Good Evening !");
	}

	//using defer
	returnedValue := tryDefer();
	fmt.Println("returnedValue : ",returnedValue);
	tryMultipleDeferWow();

	fmt.Println("===============");
}

//NOTE : defer is stack based , at it's local function
func tryDefer()(string){	//defer only defered at the end of the function
	defer fmt.Println(" Defer :)");
	fmt.Print("Try");

	return "TryDefer return value";
}

func tryMultipleDeferWow(){
	fmt.Println("counting");
	for i:=0; i < 10;i++{
		defer fmt.Println(i);
	}
	fmt.Println("done");
}


func sqrt(x float64) string{
	if x < 0{
		return sqrt(-x) + "i";
	}
	return fmt.Sprint(math.Sqrt(x));
}

func powWithMax(x,n,lim float64)float64{
	//declaration and if at the same time
	if v:= math.Pow(x,n); v < lim{
		return v;
	}else{
		fmt.Printf("%g >= %g\n", v, lim);
	}	//v scope is only at end of else

	return lim;
}

//complex data type zone;
//NOTE :say hola to Vertex
type Vertex struct{
	X int
	Y int
}
type Vertex2 struct{
	X, Y int
}

func testComplexDataType(){
	fmt.Println("=================");
	//zero based value  is <nil>
	var p * int;
	i := 42;
	p = &i;
	fmt.Println(" based value : ",i, *p);
	*p = 100;
	fmt.Println(" based value : ",i, *p)

	//struct here
	vertex1 := Vertex{1,2};
	pointerToVertex := &vertex1;
	fmt.Println(vertex1, vertex1.X, vertex1.Y);
	fmt.Println("access throught  Vertex y , x =",pointerToVertex.Y,pointerToVertex.X);

	vertex2_1 := Vertex2{X : 10};	//Y:0 is implicit
	vertex2_2 := Vertex2{Y : 10};
	vertex2_3 := Vertex2{};		//this is save because of zero based variable
	vertex2_4 := &Vertex2{1,2}; //has type *Vertex
	fmt.Println(vertex2_1,vertex2_2,vertex2_3,vertex2_4);

	//array
	var a [2]string;
	a[0] = "Hello";
	a[1] = "World";
	fmt.Println(a[0],a[1]);

	//array of vertext
	var arrayOfVertex2_1 = [2]Vertex2{
		Vertex2{1,2},
		Vertex2{3,1},
	};
	fmt.Println(arrayOfVertex2_1);

	primes := [6]int{2,3,5,7,11,13};
	fmt.Println("primes value : ",primes[0:4]);	//A slice does not store any data, it just describes a section of an underlying array

	//slice : return sliced array from its origin
	slicedPrimes1 := primes[0:5];
	slicedPrimes1[0] = 10;
	//sliced array is 0 based
	fmt.Println(slicedPrimes1,primes);
	fmt.Println("sliced : \n",
		slicedPrimes1[:],"\n",
		slicedPrimes1[:2],"\n",
		slicedPrimes1[2:],"\n",
	)

	//Slice literals
	intSliceLiterals := []int{1,2,3,4,5,6};
	boolSliceLiterals := []bool{true,false,true,true,false,true};
	arrayOfPrimeNumbers := []struct{
		anInt int
		aBool bool
	}{
		//{1},	<= when you using Slice Literals can't relly on Struct Contructor
		{2,false},
		{3,true},
		{4,false},
		{5,true},
	}
	fmt.Println(intSliceLiterals,boolSliceLiterals,arrayOfPrimeNumbers);
	slicedArray101();

	fmt.Println("=================");
}

func slicedArray101(){
	//sliced array range can be extended and drop the ranges
	primeNumbers := []int{2,3,5,7,11,13};
	printSlice(primeNumbers);

	primeNumbers = primeNumbers[:0];
	printSlice(primeNumbers);

	// Extend its length.
	primeNumbers = primeNumbers[:4];
	printSlice(primeNumbers);

	// Drop its first two values.
	primeNumbers = primeNumbers[2:];
	printSlice(primeNumbers);

	primeNumbers = primeNumbers[1:];
	printSlice(primeNumbers);

	var emptyArray []int;
	if(emptyArray == nil){
		fmt.Println("empty Array : ",emptyArray);
	}
	if emptyArray = []int{}; cap(emptyArray) == 0{
		fmt.Print("this array is empty : ");
		printSlice(emptyArray)
	}

	//we need make to create new sliced array
	newArray1_1 := make([]int,5);
	printSliceWithString("newArray1_1",newArray1_1);

	newArray1_1 = make([]int,0,5);
	printSliceWithString("newArray1_1",newArray1_1);

	//why CAP is still the same ? : i think for efficiency because the first element it still the same until which are desired BOUND
	newArray1_1 = newArray1_1[:2];
	printSliceWithString("newArray1_1",newArray1_1);

	newArray1_1 = newArray1_1[3:5];
	printSliceWithString("newArray1_1",newArray1_1);

	newArray1_2 := newArray1_1[0:];
	newArray1_2[0] = 100;
	printSliceWithString("newArray1_2",newArray1_2);

	//array 2 dimention
	board := [][]string{
		[]string{"_","_","_"},
		[]string{"_","_","_"},
		[]string{"_","_","_"},
	}
	board[0][0] = "X";
	board[0][2] = "O";
	board[1][1] = "O";
	board[2][0] = "X";

	//we need 1 loop array because of efficiency , we can just concat that , instead of print it one by one
	fmt.Println("==================\nsimple tic tac toe :p\n=================");
	for i := 0; i < len(board);i++{
		fmt.Printf("%s\n",strings.Join(board[i]," "));
	}
	fmt.Println("==================");

	var emptyArray2 []int;
	emptyArray2 = append(emptyArray,10);	//append is like insert from bottom in array and increase array capacity
	printSliceWithString("emptyArray2 : ",emptyArray2);

	emptyArray2 = append(emptyArray,20,30,40,50);
	printSliceWithString("emptyArray2 : ",emptyArray2);

	//NOTE : SLICE & ARRAY ARE DIFFERENT IN GOLANG
	aRealArray := [5]int{1,2,3,4,5};
	fmt.Println(aRealArray);
	aRealSlice := []int{1,2,3,4,5};
	printSlice(aRealSlice);

	//why using range to iterates over a slice or map.
	arrayPowOf2 := []int{1,2,4,8,16,32,64,128,256,512,1024};
	for i , v:= range  arrayPowOf2{
		fmt.Printf("2**%d = %d\n",i,v);
	}

	fmt.Println("===================");
	//why using _ = to skip the index
	for _, v := range arrayPowOf2{
		fmt.Printf("%d\n",v);
	}
	fmt.Println("===================");
	tryUsingMap();
	tryFunctionValue();
}

type LocalPosition struct{
	Lat, Long float64
}

func tryUsingMap(){
	var m map[string]LocalPosition;
	if(m == nil){
		fmt.Println("m is nil , value : ",m);
	}
	m = make(map[string]LocalPosition);
	m["Bell Labs"] = LocalPosition{
		400.123,
		500.558,
	}
	m["Ford Foundation"] = LocalPosition{
		800.123,
		911.558,
	}
	fmt.Println("map value : ",m)
	printMapOfLocalPosition(m);

	var mapLiteralOfLocalPosition = map[string]LocalPosition{
		"Edison Labs": LocalPosition{
			100.1,
			200.2,
		},
		"Walt Disney" :{		//trim the struct declaration
			500.5,
			600.6,
		},
	}
	printMapOfLocalPosition(mapLiteralOfLocalPosition);

	//map of map value
	mapOfMap := make(map[string]map[string]string);
	mapOfMap["alexander graham bell"] = make(map[string]string);
	mapOfMap["alexander graham bell"] ["birthday"] = "March 3, 1847";
	fmt.Println(mapOfMap);

	mapLiteralOfInt2 := make(map[string]int);
	mapLiteralOfInt2["one"] = 1;
	mapLiteralOfInt2["two"] = 2;
	fmt.Println("value of one : ",mapLiteralOfInt2["one"]);

	delete(mapLiteralOfInt2,"one");
	valueOfOne, isAvailable := mapLiteralOfInt2["one"];
	fmt.Println("is one available :", isAvailable,",value of one : ",valueOfOne);
}

func computeWithMathYeah(fn func(float64,float64)float64)float64{
	return fn(3,4);
}

func tryFunctionValue(){
	//save function at variable :D
	hypot := func(x, y float64) float64{
		return math.Sqrt(x*x + y*y);
	}
	fmt.Println("hypot value : ",hypot(3,4));	//Hypotenuse
	fmt.Println("hypot value : ",computeWithMathYeah(hypot));
	fmt.Println("pow value 2 pow 4: ",computeWithMathYeah(math.Pow));

	//this adder play with closure
	adder := func()func(int)int{
		sum := 0;
		return func(x int)int{
			sum += x;
			return sum;
		}
	}

	pos, neg := adder(), adder();
	for i :=0 ;i < 10;i++{
		fmt.Println(pos(i), neg(-2*i));
	}
}

func printMapOfLocalPosition(m map[string]LocalPosition){
	for i, v := range m{
		fmt.Printf("index : %s, value : %f , %f\n",i,v.Lat, v.Long);
	}
}

func printSliceWithString(s string,intArray []int){
	fmt.Printf("%s , len=%d cap=%d %v\n",s,len(intArray),cap(intArray),intArray);
}

func printSlice(s []int){
	fmt.Printf("len=%d cap=%d %v\n",len(s),cap(s),s);
}

/**
	land of method and interfaces
 */
type Person struct{		//the type is struct
	name string
	age int
}
//why is this called method : because it has special receiver argument.
func(p Person)getName()string{
	return p.name;
}

type MyFloat float64;
func(f MyFloat)Abs()float64{
	if f < 0{
		return float64(-f);
	}
	return float64(f);
}

type SimpleVertex struct{
	X, Y float64;
}
func(v *SimpleVertex)Abs()float64{
	return math.Sqrt(v.X * v.X + v.Y* v.Y);
}

func(v *SimpleVertex)Scale(f float64){
	v.X = v.X * f;
	v.Y = v.Y * f;
}
func scaleFunc(v *SimpleVertex, f float64){
	v.X = v.X * f;
	v.Y = v.Y * f;
}

type SimpleInterface interface {
	aFunction()
}

type SimpleStruct struct {
	simpleString string;
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func(simpleStruct SimpleStruct)aFunction(){
	fmt.Println(simpleStruct.simpleString);
}

type geometry interface{
	area() float64;
	perim() float64;
}

type rect struct{
	width, height float64;
}

type circle struct{
	radius float64;
}

func(r rect)area()float64{
	return r.width * r.height;
}
func(r rect)perim()float64{
	return r.width * r.height;
}
func(c circle)area() float64{
	return math.Pi * c.radius * c.radius
}
func(c circle)perim()float64{
	return 2 * math.Pi * c.radius;
}

func measure(g geometry){
	fmt.Println("geometry value : ",g);
	fmt.Printf("%v %T\n",g,g);
	fmt.Println(g.area());
	fmt.Println(g.perim());
}

type simpleInterface2 interface{
	aFunc()
}

type simpleStruct2 struct {
	aString string;
}

type TryEmptyInterface interface{
	EmptyInterfacefunc()
}

func(simpleClass2 *simpleStruct2)aFunc(){
	if simpleClass2 == nil{
		fmt.Println("<nil>");
		return;
	}
	fmt.Println(simpleClass2.aString);
}

func describeSimpleInterface2(i simpleInterface2){
	fmt.Printf("(%v %T)\n",i,i);
}

func testMethodAndInterfaces(){
	fmt.Println("================================")
	person1 := Person{"anon1",25};
	fmt.Println("person1 :",person1, " name of person1 :",person1.getName());

	/**
		IMPORTANT NOTE :
		https://tour.golang.org/methods/3
	*/
	customFloat1 := MyFloat(-math.Sqrt2);
	fmt.Println("customFloat1 : ",customFloat1.Abs());

	simpleVertex1 := SimpleVertex{3,4};
	simpleVertex1.Scale(10);
	fmt.Println("simple vertex 1 : ",simpleVertex1.Abs());

	//instantiate by refences
	simpleVertex2 := &SimpleVertex{3,4};
	scaleFunc(simpleVertex2,10);
	fmt.Println("simple vertex 2 : ",simpleVertex2.Abs());

	var simpleStruct1 SimpleInterface = SimpleStruct{"Hello There"};
	simpleStruct1.aFunction();

	rect1 := rect{10,10};
	circle1 := circle{10};
	measure(rect1);
	measure(circle1);

	var simpleInterface2Obj simpleInterface2;
	var simpleStruct2Obj *simpleStruct2;
	simpleInterface2Obj = simpleStruct2Obj;
	simpleInterface2Obj.aFunc();
	describeSimpleInterface2(simpleInterface2Obj);

	simpleInterface2Obj = &simpleStruct2{"Hello There"};
	simpleInterface2Obj.aFunc();
	describeSimpleInterface2(simpleInterface2Obj);

	//nil interface value
	var emptyInterface TryEmptyInterface;
	if(emptyInterface == nil){
		fmt.Println("emptyInterface is empty :)");
	}

	var interface1 interface{};
	retrieveGenericObj(interface1);

	interface1 = 100;
	retrieveGenericObj(interface1);

	interface1 = "hello";
	retrieveGenericObj(interface1);

	tryTypeAssertions();
	fmt.Println("================================");
}

type SimplePerson struct{
	Name string;
	Age int;
}
func (sp SimplePerson)String()string{
	return fmt.Sprintf("%v (%v Years)",sp.Name,sp.Age);
}

//custom error :sweat_smile:
type MyError struct{
	When time.Time;
	What string;
}
func(e *MyError)Error()string{
	return fmt.Sprintf("at %v %s", e.When,e.What);
}

func tryTypeAssertions(){
	//type assertions
	var localStringInterface interface{} = "Hello";

	localString1 := localStringInterface.(string);		//assert the type of variable
	fmt.Println(localString1);

	localString2 , ok := localStringInterface.(string);
	fmt.Println(localString2,ok);

	//NOTE : we need the **ok** variable , to suspend panic runtime error , because
	//       If localString3 does not hold a T, the statement will trigger a panic.
	localString3, ok := localStringInterface.(float64);
	fmt.Println(localString3,ok);

	do := func(i interface{}){
		//NOTE : type switch is a construct that permits several type assertions in series.
		switch v := i.(type){
		case int:
			fmt.Printf("Twice %v is %v\n",v, v*2);
			break;
		case string:
			fmt.Printf("%q is %v bytes long\n",v,len(v));
			break;
		default:
			fmt.Printf("I don't know the type : %T\n",v);
		}
	}

	do("10");
	do(10);
	do(true);

	//try stringer method
	simplePerson1 := SimplePerson{"Andrew",25};
	simplePerson2 := SimplePerson{"Rush Skyes",18};
	fmt.Println("simple persons value :", simplePerson1.String(),simplePerson2.String());

	//try error interface
	runErrorExample := func() error{
		return &MyError{	//return an instance of MyError
			time.Now(),
			"It didn't work",
		}
	}
	if err :=runErrorExample(); err != nil{
		fmt.Println(err);
	}

	//readers
	fmt.Println("==================================");
	reader := strings.NewReader("Hello , Reader!");
	aByte := make([]byte,8);
	for{
		n, err := reader.Read(aByte);
		fmt.Printf("n = %v err = %v  aByte = %v\n",n,err,aByte);
		fmt.Printf("aByte[:n] = %q\n",aByte[:n]);
		if err == io.EOF{
			break;
		}
	}
	fmt.Println("==================================");

	anImage := image.NewRGBA(image.Rect(0,0,100,100));
	fmt.Println(anImage.Bounds());
	fmt.Println(anImage.At(0,0).RGBA());
}

//why we need empty inteface ? , because empty interface is used for
func retrieveGenericObj(i interface{}){
	fmt.Printf("can retrieve generic interface : %v %T\n",i,i);
}

func testgoRoutineAndChannels(){
	fmt.Println("Test Go Routing and Channels\n ================================");
	say := func(aString string){
		for i :=0; i < 5; i++{
			time.Sleep(100 * time.Millisecond);
			fmt.Println(aString);
		}
	}
	//A goroutine is a lightweight thread managed by the Go runtime.
	go say("world");
	say("there");
	go say("hello");

	//test channels
	sum := func(s []int, c chan int){
		sumVal := 0;
		for _, v := range s{
			sumVal += v;
		}
		c <- sumVal // send sumVal to c
	}

	//NOTE : channel synchronize and communication in the single operation (fundamental idea)
	channelAndGoRoutineExample := func(){
		arrayOfInt := []int{1,2,3,4,5,6};
		channel := make(chan int);
		go sum(arrayOfInt[:len(arrayOfInt)/2],channel);
		go sum(arrayOfInt[len(arrayOfInt)/2:],channel);
		x, y := <-channel , <-channel; //receive from channel
		fmt.Println(x, y , x + y);
	}
	channelAndGoRoutineExample();

	//in go talk 2013, this is more unlikely
	tryBufferedChannel := func(){
		ch := make(chan int,2);
		ch <- 1;
		ch <- 2;
		fmt.Println("value from buffered channel :",<-ch, <- ch);
	}
	tryBufferedChannel();

	//NOTE : https://tour.golang.org/concurrency/4
	tryRangeAndCloseWithRoutine := func(){
		fibonacci := func(n int,c chan int){
			x, y := 0, 1;
			for i := 0; i < n ; i++{
				c <- x;
				x, y = y, x+y;
			}
			close(c);
		}

		c := make(chan int,10);	//with buffer 10
		n := cap(c);	//get channel capacity
		go fibonacci(n,c);
		fmt.Println("\nvalue of fibbonaci using golang : ");
		for i := range c{
			fmt.Println(i);
		}
	}
	tryRangeAndCloseWithRoutine();

	//try select in goroutine
	trySelectIngoroutine := func(){
		selectExample1 := func(){
			c1 := make(chan string);
			c2 := make(chan string);

			go func(){
				time.Sleep(time.Second * 1);
				c1 <- "one";
			}();
			go func(){
				time.Sleep(time.Second * 2);
				c2 <- "two";
			}();

			//using select : We’ll use select to await both of these values simultaneously
			for i:= 0; i < 2 ;i++{
				select {
				case msg1 := <- c1:
					fmt.Println("received",msg1);
				case msg2 := <- c2:
					fmt.Println("received",msg2);
				}
			}
		}

		selectExample2 := func(){
			
		}

		selectExample1();
	}
	trySelectIngoroutine();


	fmt.Println("================================");
}

func main() {
	testDeclaringVariable();
	testFlowControl();
	testComplexDataType();
	testMethodAndInterfaces();
	testgoRoutineAndChannels();
	fmt.Println("My Favourite number is : ", rand.Intn(10))
	fmt.Printf("Hello World\n");
	fmt.Println("Math.Pi number : ",math.Pi);
	fmt.Println("10 + 20 = ",add(10,20))
	fmt.Println("10 + 20 = ",add2(10,30))

	var1, var2 := swap("world","Hello,");		// := short variable declaration , and this can do redeclaration
	fmt.Println(var1,var2);

	fmt.Println(split(10));
}
