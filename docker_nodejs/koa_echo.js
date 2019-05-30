const Koa = require("koa");
const Router = require("koa-router");
const app = new Koa();
const router = new Router();
const mongoose = require("mongoose");
mongoose.connect(
  "mongodb://root:example@mongo:27017/test?authSource=admin",
  { useNewUrlParser: true }
);
const Schema = mongoose.Schema;
let model = mongoose.model("items", new Schema({}, { collection: "items" }));

// function sleep(ms) {
//   return new Promise(resolve => setTimeout(resolve, ms));
// }

router.get("/", async (ctx, next) => {
  // ctx.body = await model.findOne({ id: "5" }).select("-_id");
  ctx.body = await model.find({}).select("-_id");
});

app.use(router.routes()).use(router.allowedMethods());

app.listen(3000);
