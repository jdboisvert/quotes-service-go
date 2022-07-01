# quotes-service-go
A simple Go CRUD API for Quotes.

#### API Documentation
The API supports the following requests:
-   Create quote
    -   Method: POST
    -   URL: /api/quote
    -   Parameters:
        -   quote: Quote in quection (ex: 'Live life')(required)
        -   author_name: Person who the quote belongs to (ex: 'Morgan Freeman') (required)
    -   Responses:
        -   201: Quote created successfully


-   Get a list of all the quotes
    -   Method: GET
    -   URL: /api/quote
    -   Responses:
        -   200: Get all quotes successfully


-   Read details of a quote
    -   Method: GET
    -   URL: /api/quote/{id}
    -   {id}: The id of the quote in question
    -   Responses:
        -   200: Got quote successfully
        -   404: Quote does not exist


-   Update details of a quote

    -   Method: PUT
    -   URL: /api/quote/{id}
    -   {id}: The id of the quote in question
    -   Parameters:
        -   quote: Quote in question (ex: 'Live life')(required if author_name not given)
        -   author_name: Person who the quote belongs to (ex: 'Morgan Freeman') (required if quote not given)
    -   Responses:
        -   200: Quote updated successfully
            -   quote: holding details of the quote now updated
        -   404: Quote does not exist


-   Delete a quote
    -   Method: DELETE
    -   URL: /api/quote/{id}
    -   {id}: The id of the quote in question
    -   Responses:
        -   200: Quote deleted successfully
        -   404: Quote does not exist