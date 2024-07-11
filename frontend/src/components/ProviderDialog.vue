<template>
  <el-dialog title="连接数据库" v-model="visible" :close-on-click-modal="false" width="600px">
    <el-tabs tab-position="left">
      <el-tab-pane label="常规配置">
        <el-form label-width="80px">
          <el-form-item label="数据源">
            <el-select v-model="form.driver" placeholder="请选择数据源">
              <el-option label="mysql" value="mysql"></el-option>
              <el-option label="mssql" value="mssql"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="主机">
            <el-input v-model="form.host" placeholder="服务器IP:端口" clearable></el-input>
          </el-form-item>
          <el-form-item label="帐号">
            <el-input v-model="form.username" placeholder="连接数据库的帐号" clearable></el-input>
          </el-form-item>
          <el-form-item label="密码">
            <el-input type="password" v-model="form.password" placeholder="连接数据库的密码" clearable></el-input>
          </el-form-item>
          <el-form-item label="数据库">
            <el-input v-model="form.database" placeholder="数据库名称" clearable></el-input>
          </el-form-item>
        </el-form>
      </el-tab-pane>
      <el-tab-pane label="高级配置">
        <el-form label-position="top">
          <el-form-item label="需要过滤的数据库表前缀">
            <el-input v-model="form.filteredTablePrefixes" placeholder="需要过滤的表前缀，多个前缀以逗号分隔。" clearable></el-input>
          </el-form-item>
          <el-form-item label="创建时需要过滤的数据库字段">
            <el-input v-model="form.filteredCreatedColumns" placeholder="需要过滤的创建字段，多个字段以逗号分隔。" clearable></el-input>
          </el-form-item>
          <el-form-item label="更新时需要过滤的数据库字段">
            <el-input v-model="form.filteredUpdatedColumns" placeholder="需要过滤的更新字段，多个字段以逗号分隔。" clearable></el-input>
          </el-form-item>
        </el-form>
      </el-tab-pane>
    </el-tabs>
    <el-row slot="footer" style="justify-content: center; margin-top: 20px">
      <el-button type="primary" @click="connect">确定</el-button>
      <el-button @click="visible = false">取消</el-button>
    </el-row>
  </el-dialog>
</template>

<script setup>
import { ref, reactive, defineExpose } from 'vue'
import { storeToRefs } from 'pinia'
import { useStore } from '@/store/index.js'

const store = useStore()
const { provider } = storeToRefs(store)
const visible = ref(false)
const form = reactive({})

const show = () => {
  Object.assign(form, provider.value)
  visible.value = true
}

const hide = () => {
  visible.value = false
}

const connect = async () => {
  const ok = await store.connect(form)
  if (ok) {
    hide()
  }
}

defineExpose({
  show,
  hide,
})
</script>
