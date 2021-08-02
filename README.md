Q : Describe what you think happened that caused those bad reviews during our 12.12 event and why it happened. 

A: 
 I think there are many reason why this is happen. as example
 
	1. This happened because the quantity of product doesn't updated quickly after the customer's transaction
	2. There is no checking product's quantity when customer checking out
	3. Inventory product system does not integrated well with online shop system 
	4. There are some race condition on the coding side that made the application saves wrong data 
 
Q : Based on your analysis, propose a solution that will prevent the incidents from occurring again
A : 
    
    1. Validate for product's quantity when customers checking out. Make sure that the product that will be buyed by customer is still available 
	2. Make sure there is no race condition on coding side
	3. Put a "minimal stock" for every product that we have in the warehouse. if the product's quantity has reached its minimal stock we can show a 
       "please grab it fast" message to customer. so we can push the customers to pay the poduct quickly or we can say it as a "out of stock" if its has reached minimal stock.
	   Its also give us a spare from product's quantity if suddenly we got a broke product or missing product in warehouse.
	4. For every product's quantity on "check out" state we should reduce "current" product's quantity and put it on "on_hold" product's quantity. Also we must give
       a limit time for customers to pay theirs product. as soon as customer payed the products we should reduce the "on_hold" product's quantity and put in on "sold"
	   product's quantity. BUT if the customer doesn't pay the bill we must take back "on hold" product's quantity to "current" product's quantity   	
	5. Use inventory adjustment to update product's quantity insted. NEVER update product's quantity without using inventory adjustment	