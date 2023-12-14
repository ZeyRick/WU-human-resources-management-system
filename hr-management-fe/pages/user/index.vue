<template>
  <n-layout style="flex-grow: 1; display: flex; flex-direction: column">
    <n-card
      content-style="padding: 10px;"
      style="height: 50px; overflow: hidden"
    >
      <div
        style="
          flex-direction: row;
          display: flex;
          align-items: center;
          justify-content: space-between;
          overflow: hidden;
        "
      >
        <n-text type="primary" style="font-size: 18px">User Table</n-text>
        <n-button
          strong
          style="background-color: #409eff"
          color="#5cb85c"
          text-color="#000000"
          @click="openCreateModal"
        >
          <template #icon>
            <n-icon color="#000000">
              <AddCircleOutline />
            </n-icon>
          </template>
          Create
        </n-button>
      </div></n-card
    >
    <n-data-table :columns="tableColumns" :data="data" />
    <n-card
      content-style="padding: 10px;"
      style="display: flex; align-items: center; height: 50px; overflow: hidden"
    >
      <n-pagination :page="1" :page-count="100" />
    </n-card>

    <n-modal
      :show="showCreateModal"
      :mask-closable="false"
      @negative-click="closeCreateModal"
      @positive-click="onSubmitCreate"
    >
      <n-card
        style="width: 600px"
        title="Create New User"
        :bordered="false"
        size="huge"
        role="dialog"
        aria-modal="true"
      >
        <n-form
          ref="createFormRef"
          :rules="CommonFormRules"
          :model="createFormData"
        >
          <n-form-item path="userName" label="UserName">
            <n-input
              v-model:value="createFormData.userName"
              @keydown.enter.prevent
            />
          </n-form-item>
          <n-form-item path="password" label="Password">
            <n-input
              type="password"
              v-model:value="createFormData.password"
              @keydown.enter.prevent
            />
          </n-form-item>
        </n-form>
        <div style="display: flex; gap: 10px; justify-content: flex-end">
          <n-button round @click="closeCreateModal"> Cancel </n-button>
          <n-button round @click="onSubmitCreate"> Create </n-button>
        </div>
      </n-card>
    </n-modal>
  </n-layout>
</template>

<script setup lang="ts">
import { AddCircleOutline } from "@vicons/ionicons5";
import { tableColumns } from "./table-columns";
import { CommonFormRules } from "../../constants/formRules";
import { type FormInst, type FormValidationError } from "naive-ui";
import { apiCreateUser } from '../../apis/user'

const showCreateModal = ref<boolean>(false);
const createFormRef = ref<FormInst>();
interface CreateUserType {
  userName: string;
  password: String;
}
const createFormData = ref<CreateUserType>({
  userName: "",
  password: "",
});

const onSubmitCreate = () => {
  createFormRef.value?.validate(
    (errors: Array<FormValidationError> | undefined) => {
      if (!errors) {
        console.log(createFormData.value);
        apiCreateUser()
      } else {
        console.log(errors);
      }
    }
  );
};

const closeCreateModal = () => (showCreateModal.value = false);

const openCreateModal = () => (showCreateModal.value = true);

const data = Array.from({ length: 10 }).map((_, index) => ({
  name: `Edward King ${index}`,
  age: 32,
  address: `London, Park Lane no. ${index}`,
}));

definePageMeta({
  layout: "main",
});
</script>
