import camelcaseKeys from "camelcase-keys";

export class RPC {
  private prefix = "twirp";

  constructor(private url: string, public token: string) {}

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
        headers: {
          Authorization: `Bearer ${this.token}`,
          "Content-Type": "application/json",
        },
      })
        .then((res) => res.json())
        .then((data) => resolve(camelcaseKeys(data, { deep: true })))
        .catch(reject);
    });
  }
}
