import { Box, Button, Card, CardContent, Typography } from "@mui/material";
import { Add, SettingsApplications } from "@mui/icons-material";
import { useFieldArray, useFormContext, useWatch } from "react-hook-form";
import { EnvVariableEditor } from "@agent-management-platform/views";

export const EnvironmentVariable = () => {
    const { control, formState: { errors }, register } = useFormContext();
    const { fields, append, remove } = useFieldArray({ control, name: 'env' });
    const envValues = useWatch({ control, name: 'env' }) || [];

    const isOneEmpty = envValues.some((e: any) => !e?.key || !e?.value);

    return (
        <Card variant="outlined">
            <CardContent>
                <Box display="flex" flexDirection="row" alignItems="center" gap={1}>
                    <SettingsApplications fontSize="medium" color="disabled" />
                    <Typography variant="h5">
                        Environment Variables (Optional)
                    </Typography>
                </Box>
                <Typography variant="body2" color="text.secondary">
                    Set environment variables for your agent.
                </Typography>
                <Box display="flex" flexDirection="column" pt={2} gap={2}>
                    {fields.map((field: any, index: number) => (
                        <EnvVariableEditor
                            key={field.id}
                            fieldName="env"
                            index={index}
                            fieldId={field.id}
                            register={register}
                            errors={errors}
                            onRemove={() => remove(index)}
                        />
                    ))}
                </Box>
                <Button startIcon={<Add fontSize="small" />} disabled={isOneEmpty} variant="outlined" color="primary" onClick={() => append({ key: '', value: '' })}>
                    Add Environment Variable
                </Button>
            </CardContent>
        </Card>
    );
};
