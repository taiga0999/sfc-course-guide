import axios from 'axios';

export default{
  install(vue) {
    const idGenerator = (function* () {
      let id = 0;
      while (true) yield id++;
    }());

    vue.prototype.requests = [];

    vue.prototype.Request = class {
      constructor({
        method, url, data, config,
      }) {
        this.id = idGenerator.next().value;

        method && (this.method = method.toLowerCase());
        url && (this.url = url);
        data && (this.data = data);
        config && (this.config = config);

        this.source = axios.CancelToken.source();
        this.timestamp = Number(new Date());
      }

      perform({ onThen, onCatch, onFinally }) {
        onThen && (this.resolve = onThen);
        this.reject = onCatch || ((reason) => {
          if (axios.isCancel(reason)) {
            console.log(reason);
          } else {
            console.error(reason);
          }
        });
        onFinally && (this.onFinally = onFinally);

        vue.prototype.requests.push(this);
        axios({
          method: this.method,
          url: this.url,
          cancelToken: this.source.token,
          data: this.data,
          // config: this.config,
        })
          .then((response) => {
            //  Cancel request if any search request finished first
            if (vue.prototype.requests.length > 1) {
              let index = 0;
              vue.prototype.requests.every((request, idx) => {
                if (request.timestamp < this.timestamp) {
                  request.cancel();
                }
                if (request.id === this.id) {
                  index = idx;
                  return false;
                }
                return true;
              });
              vue.prototype.requests = vue.prototype.requests.slice(index);
            }

            this.resolve(response);
          })
          .catch(this.reject)
          .finally(this.onFinally);
      }

      cancel(cancelMsg) {
        this.source.cancel(cancelMsg);
      }
    };
  },
};
