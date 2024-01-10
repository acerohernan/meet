import camelcaseKeys from "camelcase-keys";

export class RPC {
  private prefix = "twirp";
  private headers: Record<string, string> = {
    "Content-Type": "application/json",
  };

  constructor(private url: string, private token: string) {
    if (token) {
      this.headers["Authorization"] = `Bearer ${this.token}`;
    }
  }

  async request(
    service: string,
    func: string,
    body: any,
    version = "v1"
  ): Promise<any> {
    const url = `${this.url}/${this.prefix}/${this.prefix}.${version}.${service}/${func}`;

    return new Promise((resolve, reject) => {
      fetch(url, {
        method: "POST",
        body: JSON.stringify(body),
        headers: this.headers,
      })
        .then((res) => res.json())
        .then((data) => resolve(camelcaseKeys(data, { deep: true })))
        .catch(reject);
    });
  }
}
