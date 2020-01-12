// This is a basic Flutter widget test.
//
// To perform an interaction with a widget in your test, use the WidgetTester
// utility that Flutter provides. For example, you can send tap and scroll
// gestures. You can also use WidgetTester to find child widgets in the widget
// tree, read text, and verify that the values of widget properties are correct.

import 'package:flutter/material.dart';


//----data
class User  {
  final String name;
  final String description;
  bool isSelect = false;
  User(this.name, this.description);
  IconData icons() {
    if (isSelect) {
       return  Icons.fastfood;
    }else{
      return Icons.ac_unit;
    }
  }
  bool IsSelect(){
    return isSelect;
  }
  void setSelect(){
    isSelect = !isSelect;
  }
}

//----search
class Search {
  var searchInfo = "";
  var isSearch = false;
  
}

UserMgr user = new UserMgr();

class UserMgr {
  List<User> users = new List.generate(50, (i) => new User('xiaobai $i', 'index $i'));
  var userChoice = List();
  void addChoice(int index){
    userChoice.add(index);
  }
  void delChoice(int index){
    userChoice.remove(index);
  }
  int getChoiceLenth(){
    print("lenth");
    print(userChoice.length);
     return userChoice.length;
  }
}
//--main
void main() {
  runApp(
      new MaterialApp(
        title: 'list',
        home: listWidge(),
  ));
}

//--vertical and horizontal
class listWidge extends StatelessWidget {
  //vertical
  @override
  Widget build(BuildContext context) {
    final divider =  Divider(
      height: 2.0,
      indent: 0.0,
      color: const Color(0xFF2399ff),
    );
     return MaterialApp (
        title: 'V + H',
        home: Scaffold(
          appBar: AppBar(
            title: Text('v + h + s test'),
          ),
          body: ListView(
              children: <Widget>[
                Container(
                  height: 100,
                  child:UserSearchWidge() ,
                ),
                divider,
                Container(
                  height: 100,
                  child:UserChoiceListWidge() ,
                ),
                divider,
                Container(
                  height: 600,
                  child:UserListWidge() ,
                ),
                ]
          )
        ),
     );
  }
}

//--search
class UserSearchWidge extends StatefulWidget {
  @override
  State<StatefulWidget> createState() => sHandle;
}

final UserSearch sHandle = new UserSearch();

class UserSearch extends State<UserSearchWidge> {
  @override
  Widget build(BuildContext context) {
    if (user.getChoiceLenth() )
    return Scaffold(
      body: Container(
        margin: const EdgeInsets.only(top: 25.0),
        child: Row(
          children: <Widget>[
            //--back
            Container(
              height: 45.0,
              child: Icon(
                Icons.chevron_left,
                color: Colors.black,
              ),
            ),
        Container(
          height: 45.0,
          width: 100,
          child: TextField(
            autofocus: true,
            style: TextStyle(
              color: Colors.black38,
              fontSize: 16,
            ),
          ),
        ),

            Container(
              height: 45.0,
              child: Icon(
                Icons.search,
                color: Colors.black,
              ),
            ),

           //search
          ],
        ),
      )
    );
  }
}
//--vertical
class UserListWidge extends StatefulWidget {
  @override
  State<StatefulWidget> createState() => vHandle;
}

final VUserList vHandle = new VUserList();

class VUserList extends State<UserListWidge>  {
  void changeLeading(int index) {
    setState(() {
      print("setState");
      user.users[index].setSelect();
      if (user.users[index].IsSelect() == true){
        user.addChoice(index);
      }else{
        user.delChoice(index);
      }
    });
  }
  void changeLeaded() {
    setState(() {
      print("setState2");
    });
  }
  @override
  Widget build(BuildContext context) {
    print("setState3");
    return  new Scaffold(
      body:  new ListView.builder(
        itemCount: user.users.length,
        itemBuilder: (context, index) {
          return  new ListTile(
            title: new Text(user.users[index].name),
            leading:  new Icon( user.users[index].icons()),
            onTap: (){
              this.changeLeading(index);
              hHandle.changeLeaded();
            },
           );
        },
          )
      );
  }
}

//--horizontal
class UserChoiceListWidge extends StatefulWidget {
  @override
  State<StatefulWidget> createState() => hHandle;
}


final HUserList hHandle = new HUserList();

class HUserList extends State<UserChoiceListWidge>  {
  void changeLeading(int index) {
    setState(() {
      var userIndex =  user.userChoice[index];
      print("setState");
      print(user.users[userIndex].name);
      user.delChoice(userIndex);
      user.users[userIndex].setSelect();
    });
  }
  void changeLeaded() {
    setState(() {
      print("setState2");
    });
  }
  @override
  Widget build(BuildContext context) {
    print("setState3");
    return  new Scaffold(
        body:  new ListView.builder(
          scrollDirection: Axis.horizontal,
          itemCount: user.getChoiceLenth(),
          itemBuilder: (context, index) {
            return Container(color: Colors.white , width: 80,height: 40,
                child:  GestureDetector(
                  onTap: (){
                    this.changeLeading(index);
                    vHandle.changeLeaded();
                  },
                  child: Icon(user.users[user.userChoice[index]].icons()),
                )
            );
          },
        )
    );
  }
}