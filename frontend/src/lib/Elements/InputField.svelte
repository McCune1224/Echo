<script lang="ts">
    var labelStyle = "block mb-2 text-4lx font-medium text-white dark:text-white-500";
    var inputStyle = "bg-gray-50 p-4 border border-gray-500 text-gray-500 text-5lx rounded-lg block w-full drop-shadow-shine-empty " +
    "focus:ring-white-500 focus:border-white-500 transition-fast focus:drop-shadow-shine-white " + 
    "dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-white-500 dark:focus:border-white-500";
    var msgStyle = "mt-2 text-sm text-600 dark:text-500";

    export var label;
    export var placeholder = "";
    export var message = "";

    export var requiredField = Boolean(false);

    export var type = "text";
    export var callback; //Make sure to have a parameter
</script>

<div class="mb-2 w-full">
    <label class={labelStyle} for="input">{label}</label> 
        <input class={inputStyle} on:change={(event) => callback(event.currentTarget.value)} type={type} placeholder={placeholder} required={requiredField}/> 
    <p class={msgStyle}><span class="font-medium">{message}</span></p>
</div>