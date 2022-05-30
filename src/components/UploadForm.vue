<template>
  <div>
    <single-file-upload class="upload" v-model:is-submitted="isSubmitted"
                        @get-file-id="getFileID"
                        :file-u-r-l="fileURL"/>
    <el-form :model="form" label-width="15em" label-position="right" v-if="!isSubmitted">
      <el-form-item label="Protected by password">
        <el-switch v-model="form.passwd.enable"></el-switch>
      </el-form-item>
      <el-form-item label="Enter your password" v-if="form.passwd.enable">
      <el-input placeholder="Please input password" v-model="form.passwd.ctx"
                show-password></el-input>
      </el-form-item>
      <el-form-item label="Auto Destroy">
        <el-switch v-model="form.autoDestroy.enable"></el-switch>
      </el-form-item>
      <el-form-item label="File will be deleted on" v-if="form.autoDestroy.enable">
      <el-date-picker
        v-model="form.autoDestroy.ctx"
        type="datetime"
        placeholder="Select date and time">
      </el-date-picker>
      </el-form-item>
      <el-form-item>
      </el-form-item>
    </el-form>
    <el-button type="primary" @click="isSubmitted = true">Submit</el-button>
  </div>
</template>

<script>
import SingleFileUpload from '@/components/SingleFileUpload.vue';

export default {
  name: 'UploadForm',
  components: {
    SingleFileUpload,
  },
  data() {
    return {
      isSubmitted: false,
      fileURL: '',
      form: {
        passwd: {
          enable: false,
          ctx: '',
        },
        autoDestroy: {
          enable: false,
          ctx: '',
        },
      },
    };
  },
  methods: {
    async getFileID(id) {
      fetch(`${window.location.origin}/api/createLink`, {
        method: 'POST',
        body: JSON.stringify({
          id,
          tkn: this.form.passwd.ctx,
          burn_date: this.form.autoDestroy.ctx,
        }),
      }).then((resp) => resp.json()).then((data) => {
        this.fileURL = `${window.location.origin}/file/${data.id}`;
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

.upload {
  margin: 10em 0 5em 0;
}

.el-button {
  min-width: 13.750em;
  max-width: 31.250em;
}
</style>
