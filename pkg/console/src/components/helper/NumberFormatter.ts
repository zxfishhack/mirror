import Component from "vue-class-component";
import Vue from "vue";
const humanize = require("humanize");
@Component
export default class NumberFormatter extends Vue {
  n(v?: number, decimals = 2, decPoint = ".", thousandsSep = ","): string {
    if (v) {
      return humanize.numberFormat(v, decimals, decPoint, thousandsSep);
    }
    return "-";
  }
}
