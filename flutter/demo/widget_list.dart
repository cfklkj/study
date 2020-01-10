
// This is a basic Flutter widget test.

//

// To perform an interaction with a widget in your test, use the WidgetTester

// utility that Flutter provides. For example, you can send tap and scroll

// gestures. You can also use WidgetTester to find child widgets in the widget

// tree, read text, and verify that the values of widget properties are correct.

 

import 'package:flutter/material.dart';

 

 

typedef void CartChangedCallback(User user, bool icCart);

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

 

List<User> users = new List.generate(50, (i) => new User('xiaobai $i', 'index $i'));

 

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

     return MaterialApp (

        title: 'V + H',

        home: Scaffold(

          appBar: AppBar(

            title: Text('v + h test'),

          ),

          body: ListView(

              children: <Widget>[

                Container(

                  height: 100,

                  child:UserChoiceListWidge() ,

                ),

                Divider(

                  height: 2.0,

                  indent: 0.0,

                  color: const Color(0xFF2399ff),

                ),

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

      print(users[index].name);

      users[index].setSelect();

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

        itemCount: users.length,

        itemBuilder: (context, index) {

          return  new ListTile(

             title: new Text(users[index].name),

             leading:  new Icon( users[index].icons()),

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

      print("setState");

      print(users[index].name);

      users[index].setSelect();

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

          itemCount: users.length,

          itemBuilder: (context, index) {

            return Container(color: Colors.white ,margin: EdgeInsets.only(top: 10),width: 80,height: 40,

                child:  GestureDetector(

                  onTap: (){

                    this.changeLeading(index);

                    vHandle.changeLeaded();

                  },

                  child: Icon(users[index].icons()),

                )

            );

          },

        )

    );

  }

}
