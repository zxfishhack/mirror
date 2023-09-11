import Server from "./server";
const SERVICE_NAME = process.env.VUE_APP_PORTAL;
export default {
  getRules(): Promise<any[]> {
    return Server.get("/rules");
  },
  deleteRule(id: string) {
    return Server.delete(`/rule/${id}`);
  },
  addRule(rule: Rule) {
    return Server.post("/rule", rule);
  },
  editRule(rule: Rule) {
    return Server.post(`/rule/${rule.id}`, rule);
  },
  activeRules() {
    return Server.get("/reconfigure");
  },
};

export interface ServerInterface {
  [index: string]: any;
}
