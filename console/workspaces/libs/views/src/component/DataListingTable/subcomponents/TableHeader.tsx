import { TableHead, TableRow, TableCell, TableSortLabel, useTheme } from '@mui/material';
import { ArrowUpward, ArrowDownward, TaskOutlined } from '@mui/icons-material';
import { TableColumn } from '../DataListingTable';

export interface TableHeaderProps<T = any> {
  columns: TableColumn<T>[];
  sortBy: keyof T | string;
  sortDirection: 'asc' | 'desc';
  onSort: (columnId: keyof T | string) => void;
  hasActions?: boolean;
}

export const TableHeader = <T extends Record<string, any>>({
  columns,
  sortBy,
  sortDirection,
  onSort,
  hasActions = false,
}: TableHeaderProps<T>) => {
  const theme = useTheme();

  const renderSortIcon = (columnId: keyof T | string) => {
    if (sortBy !== columnId) return null;
    return sortDirection === 'asc' ? <ArrowUpward fontSize="small" /> : <ArrowDownward fontSize="small" />;
  };

  return (
    <TableHead>
      <TableRow>
        {columns.map((column) => (
          <TableCell
            key={String(column.id)}
            align={column.align || 'left'}
            sx={{ 
              width: column.width,
              // fontWeight: 600,
              backgroundColor: 'transparent',
              borderBottom: 'none',
              padding: theme.spacing(1.5, 2),
              // fontSize: '0.875rem',
              fontSize: theme.typography.caption.fontSize,
              color: theme.palette.text.secondary,
            }}
          >
            {column.sortable !== false ? (
              <TableSortLabel
                active={sortBy === column.id}
                direction={sortBy === column.id ? sortDirection : 'asc'}
                onClick={() => onSort(column.id)}
                IconComponent={() => renderSortIcon(column.id)}
                sx={{
                  '&.MuiTableSortLabel-root': {
                    color: theme.palette.text.secondary,
                    '&:hover': {
                      color: theme.palette.text.primary,
                    },
                    '&.Mui-active': {
                      color: theme.palette.text.primary,
                    },
                  },
                }}
              >
                {column.label}
              </TableSortLabel>
            ) : (
              column.label
            )}
          </TableCell>
        ))}
        {hasActions && (
          <TableCell align="right" sx={{ 
            // fontWeight: 600,
            backgroundColor: 'transparent',
            borderBottom: 'none',
            padding: theme.spacing(1.5, 2),
            // fontSize: '0.875rem',
            color: theme.palette.text.secondary,
          }}>
            <TaskOutlined />
          </TableCell>
        )}
      </TableRow>
    </TableHead>
  );
};
