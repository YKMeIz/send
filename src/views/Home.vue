<template>
  <div>
    <upload-form />
  </div>
</template>

<script>
// @ is an alias to /src
import UploadForm from '@/components/UploadForm.vue';

export default {
  name: 'Home',
  components: {
    UploadForm,
  },
};
</script>
