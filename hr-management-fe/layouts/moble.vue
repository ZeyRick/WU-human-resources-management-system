<template>
  <div class="outer">
  <div class="middle">
    <div class="inner">
       <slot style="flex-grow: 1;" />
    </div>
  </div>
</div>
     

</template>

<script setup lang="ts">
</script>

<style>
.outer {
  display: table;
  position: absolute;
  top: 0;
  left: 0;
  height: 100%;
  width: 100%;
}

.middle {
  display: table-cell;
  vertical-align: middle;
}

.inner {
  margin-left: auto;
  margin-right: auto;
  max-width: 480px;
  height: 100%;
  display: flex;
}
</style>
