import { Box, IconButton, TextField } from "@mui/material";
import { DeleteOutline } from "@mui/icons-material";
import { FieldErrors, UseFormRegister } from "react-hook-form";

export interface EnvVariableEditorProps {
    /**
     * The field name in the form (e.g., 'env', 'environmentVariables')
     */
    fieldName: string;
    /**
     * Index of the environment variable in the array
     */
    index: number;
    /**
     * Unique field ID from react-hook-form
     */
    fieldId: string;
    /**
     * React Hook Form register function
     */
    register: UseFormRegister<any>;
    /**
     * Form errors object
     */
    errors: FieldErrors<any>;
    /**
     * Callback to remove this environment variable
     */
    onRemove: () => void;
    /**
     * Label for the key field (default: "Key")
     */
    keyLabel?: string;
    /**
     * Label for the value field (default: "Value")
     */
    valueLabel?: string;
    /**
     * Whether the value field should be a password type (default: false)
     */
    isValueSecret?: boolean;
}

export function EnvVariableEditor({
    fieldName,
    index,
    fieldId,
    register,
    errors,
    onRemove,
    keyLabel = "Key",
    valueLabel = "Value",
    isValueSecret = false
}: EnvVariableEditorProps) {
    return (
        <Box key={fieldId} display="flex" flexDirection="row" gap={2}>
            <TextField
                label={keyLabel}
                fullWidth
                {...register(`${fieldName}.${index}.key` as const)}
                error={!!(errors as any)?.[fieldName]?.[index]?.key}
                helperText={(errors as any)?.[fieldName]?.[index]?.key?.message as string}
            />
            <TextField
                label={valueLabel}
                type={isValueSecret ? "password" : "text"}
                fullWidth
                {...register(`${fieldName}.${index}.value` as const)}
                error={!!(errors as any)?.[fieldName]?.[index]?.value}
                helperText={(errors as any)?.[fieldName]?.[index]?.value?.message as string}
            />
            <IconButton size="small" color="primary" onClick={onRemove}>
                <DeleteOutline fontSize="small" color="error" />
            </IconButton>
        </Box>
    );
}

