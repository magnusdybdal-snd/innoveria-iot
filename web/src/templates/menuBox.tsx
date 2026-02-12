import Box from "@mui/material/Box";
import Link from "@mui/material/Link";

type MenuBoxProps = {
    title: string;
    add: boolean;
};

{/*Box for menu components*/}
export default function MenuBox({ title, add }: MenuBoxProps) {
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
                <h2 className="font-normal text-3xl" style={{ margin: 0 }}>
                    {title}
                </h2>
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