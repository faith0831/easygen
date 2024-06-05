import { ElMessage } from 'element-plus'
import { ElNotification } from 'element-plus'

export default {
  error(msg) {
    ElMessage({
      message: msg,
      type: 'error',
    })
  },
  success(msg) {
    ElMessage({
      message: msg,
      type: 'success',
    })
  },
  warning(msg) {
    ElMessage({
      message: msg,
      type: 'warning',
    })
  },
  notification(options) {
    ElNotification(options)
  },
}
