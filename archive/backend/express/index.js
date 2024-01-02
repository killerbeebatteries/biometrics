const express = require("express");
const app = express();
const port = 3001;

const biometrics_model = require("./biometricsModel");

app.use(express.json());
app.use(function (req, res, next) {
  res.setHeader("Access-Control-Allow-Origin", "http://localhost:3000");
  res.setHeader("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS");
  res.setHeader(
    "Access-Control-Allow-Headers",
    "Content-Type, Access-Control-Allow-Headers",
  );
  next();
});

app.get("/", (req, res) => {
  biometrics_model.getBPAndWeight()
    .then((response) => {
      res.status(200).send(response);
    })
    .catch((error) => {
      res.status(500).send(error);
    });
});

app.post("/biometrics", (req, res) => {
  biometrics_model.createBPAndWeight(req.body)
    .then((response) => {
      res.status(200).send(response);
    })
    .catch((error) => {
      res.status(500).send(error);
    });
});

app.delete("/biometrics/:id", (req, res) => {
  biometrics_model.deleteBPAndWeight(req.params.id)
    .then((response) => {
      res.status(200).send(response);
    })
    .catch((error) => {
      res.status(500).send(error);
    });
});
app.put("/biometrics/:id", (req, res) => {
  const id = req.params.id;
  const body = req.body;
  biometrics_model
    .updateBPAndWeight(id, body)
    .then((response) => {
      res.status(200).send(response);
    })
    .catch((error) => {
      res.status(500).send(error);
    });
});

app.listen(port, () => {
  console.log(`App running on port ${port}.`);
});
