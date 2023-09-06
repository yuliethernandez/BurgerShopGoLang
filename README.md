# BurgerShopGoLang
In this activity, you will create an online ordering system for a burger shop.

The program will perform the following tasks:

Take an order from a user, including a burger, a drink, a side, or a combo.
Customize selected items, including toppings for the burger.
Display the completed order to the user.
Calculate and display the order total.
The user can cancel the order and exit the program at any prompt.
If the user enters an unexpected value, the program should provide appropriate feedback and prompt them to try again.
The program must include the following features:

Create a struct that represents a burger with the following attributes:

Bun: Boolean (True if you want the burger with bun, False otherwise)
Price: float64
Dressed: boolean (True if we want the burger fully dressed, False otherwise.)
Do a basic boolean option for the first pass. You can add other options after the main program works as expected.
Create a struct that represents a drink with the following attributes

Price: Float64
Type: String
Create a struct that represents a side with the following attributes

Price: Float64
name: String
Create a nested struct called Combo which represents a burger, drink, and side combo.

burger
side
drink
price: The price of the combo is the price of the three items with a 20% discount
Next, implement the following functions:

A function called user_input_burger that asks the user for how they want their burger and store it in a struct type.
The price of the burger is computed as follows:
Burger with bun: 5
Burger without bun: 4
Dressed has no impact on price.
A function for user_input_drink, which asks the user for their drink
Provide the user with a limited number of drink options to choose from (include 3-4 choices).
The price of water is 1.
The price of any other drink is 2.
A function for user_input_side, which asks the user for their side
Provide the user with a limited number of side options to choose from, such as fries, onion rings, a salad, and coleslaw
The price of fries is 1 dollar
The price of any other option is 2 dollars
A function for user_input_combo, which asks the user for their combo preferences.
A function called take_order_from_user
ask the user for a name for the order
repeat taking the order until the user is done
display order details
Display a thank you message
