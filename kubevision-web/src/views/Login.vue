<template>
  <v-container class="text-center">
    <v-card width="400" class="mx-auto " elevation="10">
      <v-img height="80" src="@/assets/favicon.svg" class="mt-4" />
      <v-card-title>欢迎使用 Kubevision</v-card-title>
      <v-card-text>
        <v-text-field density="compact" placeholder="请输入用户名" prepend-icon="mdi-account" v-model="auth.username">
        </v-text-field>
        <v-text-field density="compact" placeholder="请输入密码" v-model="auth.password" prepend-icon="mdi-lock"
          :append-inner-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'" :type="showPassword ? 'text' : 'password'"
          @click:append-inner="showPassword = !showPassword">
        </v-text-field>
      </v-card-text>

      <v-divider></v-divider>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="warning" rounded variant="flat" style="width: 40%;" text="登录" @click="login()">登录</v-btn>
        <v-spacer></v-spacer>
      </v-card-actions>
    </v-card>
    <v-notifications location="bottom right" :timeout="3000" />
  </v-container>
</template>

<script setup>
import { ref, getCurrentInstance } from 'vue';

import API from '@/assets/app/api';
import notify from '@/assets/app/notify';

var showPassword = ref(false);
var refreshingRegion = ref(false);

const auth = ref({ cluster: null, region: null, username: null, password: null, });
const { proxy } = getCurrentInstance()

const clusters = ref([])
const regions = ref([])

async function refreshClusters() {
  clusters.value = (await API.cluster.list()).clusters
  if (clusters.value.length > 0 && !auth.value.cluster) {
    auth.value.cluster = clusters.value[0].id
  }
  refreshRegions(true)
}

async function refreshRegions(force = false) {
  if (!force && regions.value.length > 0) {
    return
  }
  if (!auth.value.cluster) {
    return
  }
  refreshingRegion.value = true;
  regions.value = await API.cluster.regions(auth.value.cluster)
  refreshingRegion.value = false;
  if (regions.value.length > 0 && !auth.value.region) {
    auth.value.region = regions.value[0]
  }
}

async function login() {
  try {
    let resp = await API.system.login(auth.value.username, auth.value.password)
    notify.success('登录成功')
    localStorage.setItem('X-Auth-Token', resp.headers.get('X-Auth-Token'));
    proxy.$router.push('/dashboard')
  } catch (e) {
    console.error(e)
    notify.error('登录失败')
  }
}

// localStorage.removeItem('namespace')
// refreshClusters()

</script>
