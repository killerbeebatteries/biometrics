require("dotenv").config();

const Pool = require("pg").Pool;
const pool = new Pool({
  user: process.env.DB_USER,
  host: process.env.DB_HOST,
  database: process.env.DB_NAME,
  password: process.env.DB_PASS,
  port: 5432,
});

// get all bp and weight data
const getBPAndWeight = async () => {
  try {
    return await new Promise(function (resolve, reject) {
      pool.query("SELECT * FROM bp_and_weight", (error, results) => {
        if (error) {
          reject(error);
        }
        if (results && results.rows) {
          resolve(results.rows);
        } else {
          reject(new Error("No results found"));
        }
      });
    });
  } catch (error_1) {
    console.error(error_1);
    throw new Error("Internal server error");
  }
};

// create bp and weight data
const createBPAndWeight = async (body) => {
  return new Promise(function (resolve, reject) {
    const { date, time, sys, dia, bp, weight_total, weight_fat, weight_muscle, comment } = body;
    pool.query(
      "INSERT INTO bp_and_weight (date, time, sys, dia, bp, weight_total, weight_fat, weight_muscle, comment) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *",
      [
        date,
        time,
        sys,
        dia,
        bp,
        weight_total,
        weight_fat,
        weight_muscle,
        comment,
      ],
      (error, results) => {
        if (error) {
          reject(error);
        }
        if (results && results.rows) {
          resolve(
            `A new BP and Weight record has been added: ${
              JSON.stringify(results.rows[0])
            }`,
          );
        } else {
          reject(new Error("No results found"));
        }
      },
    );
  });
};

// delete bp and weight data
const deleteBPAndWeight = async (id) => {
  return new Promise(function (resolve, reject) {
    pool.query(
      "DELETE FROM bp_and_weight WHERE id = $1",
      [id],
      (error, results) => {
        if (error) {
          reject(error);
        }
        resolve(`BP and Weight record deleted with ID: ${id}`);
      },
    );
  });
};

// update bp and weight data
const updateBPAndWeight = (id, body) => {
  return new Promise(function (resolve, reject) {
    const { date, time, sys, dia, bp, weight_total, weight_fat, weight_muscle, comment } = body;
    pool.query(
      "UPDATE bp_and_weight SET date = $2, time = $3, sys = $4, dia = $5, bp = $6, weight_total = $7, weight_fat = $8, weight_muscle = $9, comment = $10 WHERE id = $1 RETURNING *",
      [
        id,
        date,
        time,
        sys,
        dia,
        bp,
        weight_total,
        weight_fat,
        weight_muscle,
        comment,
      ],
      (error, results) => {
        if (error) {
          reject(error);
        }
        if (results && results.rows) {
          resolve(`Merchant updated: ${JSON.stringify(results.rows[0])}`);
        } else {
          reject(new Error("No results found"));
        }
      },
    );
  });
};

module.exports = {
  getBPAndWeight,
  createBPAndWeight,
  deleteBPAndWeight,
  updateBPAndWeight,
};
