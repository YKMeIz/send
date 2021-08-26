import {
  ElButton,
  ElContainer, ElDatePicker, ElFooter,
  ElForm,
  ElFormItem,
  ElInput,
  ElMain, ElMessageBox, ElPopover,
  ElProgress,
  ElSwitch,
} from 'element-plus';

export default (app) => {
  app.use(ElButton);
  app.use(ElProgress);
  app.use(ElForm);
  app.use(ElFormItem);
  app.use(ElDatePicker);
  app.use(ElInput);
  app.use(ElSwitch);
  app.use(ElContainer);
  app.use(ElMain);
  app.use(ElFooter);
  app.use(ElPopover);
  app.use(ElMessageBox);
};
