import Box from "@mui/material/Box";
import Link from "@mui/material/Link";
import { Link as RouterLink } from 'react-router';

type MenuBoxProps = {
    title: string;
    page: string;
    add: boolean;
};

{/*Box for menu components*/}
export default function MenuBox({ title, add, page }: MenuBoxProps) {
    return (
        <Box
            sx={{
                backgroundColor: 'secondary.light',
                color: 'primary.main',
                margin: 1,
                borderRadius: 3,
                width: 300
            }}
        >
            <div className="w-64 p-[10px]">
                {/*Page link*/}
                <Link
                    component={RouterLink} // use react-router's Link
                    to={page}               // absolute path
                    underline="none"
                    sx={{
                        fontSize: '1.75rem',
                        fontWeight: 'normal',
                        color: 'primary.main',
                    }}
                >
                    {title}
                </Link>
                <nav>
                    <Link href="#" underline="none">
                        | KPI dashboard
                    </Link>
                    <br />
                    {add && (
                        <Link href="#" underline="none">
                            + Add
                        </Link>
                    )}
                </nav>
            </div>
        </Box>
    );
}