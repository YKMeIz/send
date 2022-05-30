<template>
  <div>
    <el-progress type="circle" :percentage="progress" :indeterminate="true"
                 :color="colors" v-if="isSubmitted && progress !== 100"></el-progress>
    <form enctype="multipart/form-data" method="post" name="file" v-if="!isSubmitted">
      <input style="display:none;" id="form-input" type="file" name="file" required />
      <label for="form-input" class="upload-label">
        Click here to upload file.
      </label>
    </form>
    <div v-if="fileURL !== ''" class="ln">
      <el-popover
        placement="top"
        title="Copied!"
        :width="160"
        trigger="manual"
        content="File link has been copied to your clipboard."
        v-model:visible="copied"
      >
      <template #reference>
        <el-button type="text" @click="copyToClipboard">
          Your access link is {{ fileURL }}</el-button>
      </template>
      </el-popover>
      <br>
      <vue-qrcode :value="fileURL" :options="{ width: 200 }"></vue-qrcode>
    </div>
  </div>
</template>

<script>
export default {
  name: 'SingleFileUpload',
  props: {
    isSubmitted: Boolean,
    fileURL: String,
  },
  data() {
    return {
      progress: 0,
      copied: false,
      colors: [
        { color: '#f56c6c', percentage: 20 },
        { color: '#e6a23c', percentage: 40 },
        { color: '#5cb87a', percentage: 60 },
        { color: '#1989fa', percentage: 80 },
        { color: '#6f7ad3', percentage: 100 },
      ],
      xhr: new XMLHttpRequest(),
    };
  },
  watch: {
    isSubmitted() {
      if (this.isSubmitted) {
        this.submit();
      }
    },
  },
  methods: {
    copyToClipboard() {
      const el = document.createElement('textarea');
      el.value = this.fileURL;
      el.setAttribute('readonly', '');
      el.style.position = 'absolute';
      el.style.left = '-9999px';
      document.body.appendChild(el);
      el.select();
      document.execCommand('copy');
      document.body.removeChild(el);
      this.copied = true;
    },
    submit() {
      this.xhr = new XMLHttpRequest();
      const file = new FormData(document.forms.namedItem('file'));
      // listen for upload progress
      this.xhr.upload.onprogress = (event) => {
        this.progress = Math.round(100 * (event.loaded / event.total));
      };
      // handle error
      this.xhr.upload.onerror = () => {
        console.log(`Error: status code ${this.xhr.status}.`);
      };
      // upload completed successfully
      this.xhr.onload = () => {
        const data = JSON.parse(this.xhr.responseText);
        if (this.xhr.status === 200) {
          this.$emit('getFileId', data.id);
          return;
        }
        console.log(`${data.status} (${this.xhr.status}): ${data.message}`);
      };
      this.xhr.open('POST', `${window.location.origin}/api/uploadFile`);
      this.xhr.send(file);
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
.ln {
  margin-top: 5em;
}
.upload-label {
  cursor: pointer;
  font-size: 28px;
  color: #8c939d;
  text-align: center;

}
</style>
