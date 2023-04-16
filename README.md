### Objective

Implement a URL shortening service using Go and no framework.

### Brief

ShortLink is a URL shortening service where you enter a URL such as https://sample.io/library/react and it returns a short URL such as http://short.est/GeAi9K.

### Tasks

-   Implement assignment using:
    -   Language: **Go**
    -   Framework: **no framework**
    -   Two endpoints are required
        -   /encode - Encodes a URL to a shortened URL
        -   /decode - Decodes a shortened URL to its original URL.
    -   Both endpoints should return JSON
-   There is no restriction on how your encode/decode algorithm should work. You just need to make sure that a URL can be encoded to a short URL and the short URL can be decoded to the original URL. **You do not need to persist short URLs to a database. Keep them in memory.**
-   Provide detailed instructions on how to run your assignment in a separate markdown file
-   Provide API tests for both endpoints

### Evaluation Criteria

-   **Go** best practices
-   API implemented featuring a /encode and /decode endpoint
-   Show us your work through your commit history
-   Completeness: did you complete the features? Are all the tests running?
-   Correctness: does the functionality act in sensible, thought-out ways?
-   Maintainability: is it written in a clean, maintainable way?


