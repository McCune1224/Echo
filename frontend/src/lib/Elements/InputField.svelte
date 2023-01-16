<script lang="ts">
    var labelStyle = "block mb-2 text-4lx font-medium text-gray-700 dark:text-white-500";
    var inputStyle = "bg-gray-50 p-4 border border-gray-500 text-gray-900 text-3lx rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500";
    var msgStyle = "mt-2 text-sm text-600 dark:text-500";

    export var label;
    export var placeholder = "";
    export var message = "";

    export var requiredField = Boolean(false);

    export var type = "text";
    export var callback; //Make sure to have a parameter
</script>

<div class="mb-6 w-full">
    <label class={labelStyle} for="input">{label}</label>
    {#if requiredField}
        <input class={inputStyle} on:change={(event) => callback(event.currentTarget.value)} type={type} placeholder={placeholder} required/>
    {/if}
    {#if !requiredField}
        <input class={inputStyle} on:change={(event) => callback(event.currentTarget.value)} type={type} placeholder={placeholder}/>
    {/if}
    <p class={msgStyle}><span class="font-medium">{message}</span></p>
</div>