## 什么是 JSON

- JSON 指的是 JavaScript 对象表示法（**J**ava**S**cript **O**bject **N**otation）
- JSON 是轻量级的文本数据交换格式
- JSON 独立于语言：JSON 使用 Javascript语法来描述数据对象，但是 JSON 仍然独立于语言和平台。JSON 解析器和 JSON 库支持许多不同的编程语言。 目前非常多的动态（PHP，JSP，.NET）编程语言都支持JSON。
- JSON 具有自我描述性，更易理解



**JSON是干什么用的（理解）**：

是一种轻量级的数据交换格式。易于人阅读和编写。同时也易于机器解析和生成。

有了json，使用key-value 的模式来更加直观存储数据；

json也可用于前后端之间互相传递json数据。比如前端发起请求，调用接口，后端返回一串json数据,处理数据，渲染到页面上；

可以处理Javascript 和web服务器端的之间数据交换；

JSON 通常用于与服务端交换数据。



## JSON 语法规则 

JSON 语法是 JavaScript 对象表示语法的子集。 

- 数据在名称/值对中
- 数据由逗号分隔
- 大括号保存对象
- 中括号保存数组

### JSON数字 

JSON数字可以是整型或者浮点型

### JSON对象

JSON对象在大括号{}中书写：对象可以包含多个名称/值对	例：

```json
{ "name":"菜鸟教程" , "url":"www.runoob.com" }
```

这条语句等价于：

```json
name = "菜鸟教程"
url = "www.runoob.com"
```

对象可以包含多个 **key/value（键/值）**对。

**key **必须是字符串，**value** 可以是合法的 JSON 数据类型（字符串, 数字, 对象, 数组, 布尔值或 null）。

#### 访问对象值

现有一对象：

```json
var myObj, x;
myObj = { "name":"runoob", "alexa":10000, "site":null };
```

x = myObj.name;

或者   x = myObj["name"]	***(这里的引号不要忘了)***

#### 循环对象

可以使用for-in来循环对象的属性：

```json
var myObj = { "name":"runoob", "alexa":10000, "site":null };
for (x in myObj) {
    document.getElementById("demo").innerHTML += myObj[x] + "<br>";
	/*如果是document.getElementById("demo").innerHTML += x + "<br>";
		那么输出到是对象里的属性，如果用了中括号[]那么就是属性的值*/
}
```

上面的dociment.getElementById("demo").innerHTML += xxxx + "<br>"作用是 **输出**

#### 删除对象属性

使用关键字**delete**来删除JSON对象的属性；  例：

```json
delete myObj.sites.site1;
delete myObj.sites["site1"]   //删除嵌套对象里的对象
```



### JSON 数组

JSON对象在中括号中书写：数组可包含多个对象

```json
{
"sites": [
{ "name":"菜鸟教程" , "url":"www.runoob.com" }, 
{ "name":"google" , "url":"www.google.com" }, 
{ "name":"微博" , "url":"www.weibo.com" }
]
}//例子中，对象 "sites" 是包含三个对象的数组。每个对象代表一条关于某个网站（name、url）的记录。
```

**访问JavaScript对象数组中的第一项**：sites[0].name

**修改数据**：sites[0].name="another"



### JSON.parse()

JSON 通常用于与服务端交换数据。 在接收服务器数据时一般是字符串。

 我们可以使用 JSON.parse() 方法将数据转换为 JavaScript 对象。 

##### 语法

```json
JSON.parse(text[, reviver])
/*参数说明：  text：必需，一个有效的JSON字符串
			reviver：可选，一个转换结果的函数，将为对象的每个成员调用函数
```

例如从服务器接收了以下数据：

```
{ "name":"runoob", "alexa":10000, "site":"www.runoob.com" }
```

使用JSON.parse()方法处理以上数据，将其转换成JavaScript对象：

```json
var obj = JSON.parse('{ "name":"runoob", "alexa":10000, "site":"www.runoob.com" }');
```

**具体的从服务器接收JSON数据一些操作**参照[JSON.parse()链接](https://www.runoob.com/json/json-parse.html)

### JSON.stringify()

在向服务器发送数据时一般是字符串。 我们可以使用 JSON.stringify() 方法将 JavaScript 对象转换为字符串。 

**语法**：

```json
JSON.stringify(value[, replacer[, space]])
//参数说明
/*value：必需， 要转换的 JavaScript 值（通常为对象或数组）
  replacer：可选，用于转换结果的函数或数组
  space：可选，
```

例如向服务器发送以下数据：

```json
var obj = { "name":"runoob", "alexa":10000, "site":"www.runoob.com"};
```

我们使用  JSON.stringify()  方法处理以上数据，将其转换为字符串：

```json
var myJSON = JSON.stringify(obj);
//myJSON为字符串
```

将JavaScript数组转换为JSON字符串：

```json
var arr = [ "Google", "Runoob", "Taobao", "Facebook" ];
var myJSON = JSON.stringify(arr);
document.getElementById("demo").innerHTML = myJSON;
//myJSON为字符串；	将myJSON发送到服务器
```

