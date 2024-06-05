<template>
  <div id="app">
    <el-config-provider :size="size" :locale="locale">
      <router-view />
    </el-config-provider>
  </div>
</template>

<script>
import locale from 'element-plus/es/locale/lang/zh-cn'

export default {
  name: 'App',
  setup() {
    const size = 'default'
    return { size, locale }
  },
}
</script>
