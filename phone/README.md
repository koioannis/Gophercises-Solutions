
# Phone Number Normalizer ![language](https://img.shields.io/github/go-mod/go-version/koioannis/Gophercises-Solutions)


A PoC app that iterates through a database, normalizes all of the phone numbers and then deletes duplicates. 

For example, given phone numbers below
```
1234567890
123 456 7891
(123) 456 7892
(123) 456-7893
123-456-7894
123-456-7890
1234567892
(123)456-7892
```
The final numbers in the database would be:
```
1234567890
1234567891
1234567892
1234567893
1234567894
```
