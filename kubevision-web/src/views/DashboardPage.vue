<template>
  <v-app style="font-size: small;">
    <v-navigation-drawer :rail="navigation.mini" :width="ui.navigationWidth.value" :expand-on-hover="navigation.mini">
      <v-list-item title="Kubevision">
        <template v-slot:prepend>
          <v-avatar image="@/assets/favicon.svg" rounded="0"></v-avatar>
        </template>
      </v-list-item>
      <v-list rounded="xl" density='compact' class="pt-4" open-strategy="single">
        <div v-for="group in navigation.group" v-bind:key="group.name">
          <v-list-subheader class="text-primary">{{ $t(group.name) }} </v-list-subheader>
          <v-list-item v-for="(item, i) in group.items" v-bind:key="i" :title="$t(item.title)" :value="item"
            color="warning" @click="selectItem(item)" :disabled="$route.path.startsWith('/dashboard' + item.router)"
            :active="$route.path.startsWith('/dashboard' + item.router)">

            <template v-slot:prepend><v-icon :icon="item.icon"></v-icon></template>
          </v-list-item>
        </div>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar density="compact">
      <v-app-bar-nav-icon @click="navigation.mini = !navigation.mini"></v-app-bar-nav-icon>
      <v-toolbar-title class="ml-1">
        <v-select solo-inverted flat hide-details v-model="namespace" item-title="name" item-value="name"
          :items="namespaces" @update:modelValue="changeNamespace()">

          <template v-slot:prepend>命名空间:</template>
          <!-- @update:modelValue="changeRegion()" -->
        </v-select>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <btn-theme />
      <SettingSheet />
      <btn-logout />
    </v-app-bar>

    <v-main>
      <v-container fluid>
        <router-view></router-view>
      </v-container>
    </v-main>
    <v-notifications location="bottom right" :timeout="3000" />
  </v-app>
</template>

<script>

import SETTINGS from '@/assets/app/settings';

import BtnTheme from '../components/plugins/BtnTheme.vue';
import BtnHome from '../components/plugins/BtnHome.vue';
import BtnLogout from '../components/plugins/BtnLogout.vue';
import i18n from '@/assets/app/i18n';
import SettingSheet from '@/components/dashboard/SettingSheet.vue';
import { Utils } from '@/assets/app/lib';
import notify from '@/assets/app/notify';
import API from '@/assets/app/api';

const navigationGroup = [
  {
    name: 'cluster',
    items: [
      { title: 'home', icon: 'mdi-home', router: '/home' },
      { title: 'namespace', icon: 'mdi-alpha-h-circle', router: '/namespace' },
      { title: 'node', icon: 'mdi-alpha-h-circle', router: '/node' },
    ]
  },
  {
    name: 'application',
    items: [
      { title: 'workload', icon: 'mdi-alpha-w-circle', router: '/workload' },
      { title: 'pod', icon: 'mdi-alpha-p-circle', router: '/pod' },
      { title: 'service', icon: 'mdi-alpha-s-circle', router: '/service' },
    ]
  },
  {
    name: 'configCenter',
    items: [
      { title: 'configMap', icon: 'mdi-alpha-c-circle', router: '/configmap' },
      { title: 'secret', icon: 'mdi-alpha-s-circle', router: '/secret' },
    ]
  },
]

export default {
  components: {
    BtnTheme, BtnHome,
    SettingSheet,
    BtnLogout,
  },

  data: () => ({
    I18N: i18n,
    name: 'Kubevision',
    showSettingSheet: false,
    ui: {
      navigationWidth: SETTINGS.ui.getItem('navigatorWidth'),
    },
    namespace: sessionStorage.getItem('namespace') || 'default',
    navigation: {
      group: navigationGroup,
      selectedItem: navigationGroup[0].items[0].title,
      mini: false,
      drawer: true,
      itemIndex: 0,
    },
    namespaces: [],
  }),
  methods: {
    async refresh() {
      this.namespaces = (await API.namespaces.list()).namespaces;
    },
    selectItem(item, route) {
      this.navigation.selectedItem = item.title;
      Utils.setNavigationSelectedItem(item);
      if (!route) {
        this.$router.push('/dashboard' + item.router)
      }
      let selectedItem = this.getItem();
      this.navigation.itemIndex = selectedItem.index;
      // if (this.$route.path == '/dashboard' || this.$route.path != '/dashboard' + item.router) {
      //   this.$router.replace({ path: '/dashboard' + item.router });
      // }
    },
    getItem() {
      let localItem = Utils.getNavigationSelectedItem();
      if (this.$route.path == '/dashboard' && !localItem) {
        return { index: 0, item: navigationGroup[0].items[0] };
      }

      let selectedItemIndex = -1;
      for (let groupIndex in navigationGroup) {
        let group = navigationGroup[groupIndex];
        for (let itemIndx in group.items) {
          selectedItemIndex += 1;
          let item = group.items[itemIndx];
          if (this.$route.path == item.router || (localItem && localItem.router == item.router)) {
            return { index: selectedItemIndex, item: item }
          }
        }
      }
      return { index: 0, item: navigationGroup[0][0] };
    },
    initItem() {
      let selectedItem = this.getItem();
      this.navigation.itemIndex = selectedItem.index;
      if (this.$route.path == '/dashboard/server/new') {
        this.selectItem(selectedItem.item, '/dashboard/server/new');
      } else if (this.$route.path == '/dashboard/hypervisor/tenantUsage') {
        this.selectItem(selectedItem.item, '/dashboard/hypervisor');
      } else {
        this.selectItem(selectedItem.item,);
      }
    },
    getItemIndexByRoutePath(routePath) {
      let itemIndex = -1;
      for (let groupIndex in navigationGroup) {
        let group = navigationGroup[groupIndex];
        for (let index in group.items) {
          let item = group.items[index];
          itemIndex += 1;
          if (routePath == item.router) {
            return itemIndex;
          }
        }
      }
    },
    changeNamespace() {
      sessionStorage.setItem('namespace', this.namespace);
      location.reload();
    }
  },
  created() {
    if (!localStorage.getItem('X-Auth-Token')) {
      notify.error('请重新登录')
      this.$router.push('/login')
      return
    }
    let self = this;
    API.system.isLogin().then(function () {
      self.initItem();
      self.$vuetify.theme.dark = SETTINGS.ui.getItem('themeDark').value;
      self.refresh();
    }).catch((e) => {
      notify.error('请重新登录')
      self.$router.push('/login')
    })
  },
}
</script>
