const express = require("express");
const app = express();
const port = 8000;

app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.disable("etag");

app.use((req, res, next) => {
  console.log(req.method, req.path, req.headers["user-agent"]);
  res.removeHeader("x-powered-by");
  next();
});

app.get("/", (req, res) => {
  res.status(200).send("Lorem Ipsum Dollor");
});

app.get("/get", (req, res) => {
  res.status(200).json({ message: "Hello From Test API" });
});

app.post("/post", (req, res) => {
  let myJson = req.body; // your JSON

  res.status(200).send(myJson);
});

app.post("/postform", (req, res) => {
  res.status(200).send(JSON.stringify(req.body));
});

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`);
});
