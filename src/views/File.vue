<template>
  <div class="form">
    <el-form :model="form" label-width="15em" label-position="right">
      <el-form-item label="Enter your password">
        <el-input placeholder="Please input password" v-model="form.passwd"
                  show-password></el-input>
      </el-form-item>
       <el-form-item>
        <el-button type="primary" @click="download">Download</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  name: 'File',
  data() {
    return {
      fileURL: '',
      form: {
        passwd: '',
      },
    };
  },
  methods: {
    async download() {
      fetch(`${window.location.origin}/api/getFile`, {
        method: 'POST',
        body: JSON.stringify({
          id: this.$route.params.id,
          tkn: this.form.passwd,
        }),
      }).then((resp) => {
        if (resp.status !== 200) {
          throw Error(resp.statusText);
        }
        return resp.json();
      }).then((filename) => {
        fetch(`${window.location.origin}/ln/${this.$route.params.id}`, {
          method: 'GET',
          headers: new Headers({
            Authorization: this.form.passwd,
          }),
        }).then((response) => response.blob())
          .then((blob) => {
            const url = window.URL.createObjectURL(blob);
            const a = document.createElement('a');
            a.href = url;
            a.download = filename.id;
            // we need to append the element to the dom -> otherwise it will not work in firefox
            document.body.appendChild(a);
            a.click();
            a.remove(); // afterwards we remove the element again
          });
      }).catch((err) => {
        this.$confirm(`${err}.<br>Reason: Wrong file id, or password, or file may been deleted.`, 'ERROR', {
          confirmButtonText: 'OK',
          type: 'error',
          center: true,
          showCancelButton: false,
          dangerouslyUseHTMLString: true,
        }).catch(() => {});
      });
    },
  },
};
</script>

<style scoped>
.el-form {
  max-width: 31.250em;
  text-align: left;
}
.el-input {
  max-width: 13.750em;
}

.form {
  margin: 10em 0 5em 0;
}
</style>
