
// This is a basic Flutter widget test.

//

// To perform an interaction with a widget in your test, use the WidgetTester

// utility that Flutter provides. For example, you can send tap and scroll

// gestures. You can also use WidgetTester to find child widgets in the widget

// tree, read text, and verify that the values of widget properties are correct.

 

import 'package:flutter/material.dart';

 

 

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
 

List<User> users = new List.generate(50, (i) => new User('xiaobai $i', 'index $i'));
 
void main() {

  runApp(

      new MaterialApp(

        title: 'list',

        home: UserListWidge(),

  ));

}

 

class UserListWidge extends StatefulWidget {

  @override

  State<StatefulWidget> createState() => new UserList();

}

 

 

class UserList extends State<UserListWidge>  {

  void changeLeading(int index) {

    setState(() {

      print("setState");

      print(users[index].name);

      users[index].setSelect();

    });

  }

  @override

  Widget build(BuildContext context) {

    print("setState");

    return  new Scaffold(appBar:  new AppBar(

      title: new Text('test'),

    ),

      body:  new ListView.builder(

        itemCount: users.length,

        itemBuilder: (context, index) {

          return  new ListTile(

             title: new Text(users[index].name),

             leading:  new Icon( users[index].icons()),

            onTap: (){

              this.changeLeading(index);

            },

           );

        },

          )

      );

  }

}
