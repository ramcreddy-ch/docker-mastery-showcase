const express = require('express');
const app = express();
const port = process.env.PORT || 3000;

app.get('/', (req, res) => {
    res.json({
        message: "Level 1: Foundations",
        status: "Service is running in a basic Docker container.",
        tip: "This is the 'Standard' way, but not the 'Production' way yet!"
    });
});

app.listen(port, () => {
    console.log(`Basic app listening at http://localhost:${port}`);
});
