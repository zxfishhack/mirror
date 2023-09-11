interface Rule {
  id: string;
  prefix: string;
  postfix: string;
  replacePrefixWith: string;
  upstream: string;
  checkMD5: boolean;
  active: boolean;
}
