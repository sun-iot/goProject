# 类型系统

今天在看《GO语言编程》时，看到面向对象一章节，里面提到了一个我之前从未见过的的概念——类型系统。类型系统，顾名思义，就是指一个语言的类型体系结构。一个典型的类型系统通常包含如下基本内容：

> 1. 基础类型，如byte , int , bool , float ；
> 2. 复合类型，如数组，结构体，指针等；
> 3. 可以指向任意对象的类型（Any类型）；
> 4. 值语义和引用语义；
> 5. 面向对象，即所有具备面向对象特征（比如成员方法）的类型；
> 6. 接口

类型系统描岸述的是这些内容在一个语言中如何被关联。比如说，笔者接触最多的Java。书中也是用Java为例说明类型系统的。

在Java语言中，存在两套安全独立的类型系统：一套是**值类型系统**，主要是**基本类型**，例如byte , int , boolean , char , double ,float , long , short，这些类型基于值语义；一套是以Object类型为根的**对象类型系统**，这些类型可以定义成员变量和成员方法，可以有虚函数，基于引用语义，只允许**在堆上创建（通过new创建）**。Java语言中的Any类型就是整个对象类型系统的根——java.lang.Object类型，只有对象类型系统中的实例才可以被Any类型引用。值类型想要被Any类型引用，需要装箱（boxing）过程，比如 int 类型需要装箱成为 Integer类型。另外，只有对象类型系统中的类型才可以实现接口，具体方法是让该类型从要实现的接口继承。

## 为类型添加方法

在GO语言中，可以给任意类型（包含内置类型，但不包含指针类型）添加下相应的方法，例如：

```go
type Integer int
// 定义一个新的类型Integer，他和int没有本质区别，只是他为内置的int类型增加了个新的方法Less(),这样实现了Integer后，就可以让整型想一个普通的类一样使用
func (a Integer) Less (b Integer) bool {
    return a < b 
}
```

“在GO语言中，没有隐藏的this指针”这句话含义是：

> 1. 方法施加的目标（也就是“对象”）显式传递，没有被隐藏起来；
> 2. 方法施加的目标（也就是“对象”）不需要非得是指针，也不用非得叫this.
>
> ```java
> // 可以看看Java实现这同样的程序
> class Integer{
>     private int val ; 
>     public boolean Less(Integer b){
>         return this.val < b.val ;
>     }
> }
> ```
>
> 对于上面的Java代码，会有个思考，this到底从何而来？这主要是因为Integer类的Less()方法隐藏了第一个参数Integer* this .翻译成C语言代码，会更加清晰：
>
> ```c
> struct Integer{
>     int val ; 
> };
> bool Integer_Less(structInteger* this , structInteger* b){
>     return this->val < b->val
> }
> ```

## 值语义和引用语义

GO语言中的大多数类型基于值语义，包括：

> 基本类型：如byte , int , bool , float32 , float64 , string
>
> 复合类型：如数组（array)，结构体（struct），指针（pointer）







