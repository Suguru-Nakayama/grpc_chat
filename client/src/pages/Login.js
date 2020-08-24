import React, { useState } from 'react';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import { auth } from '../Firebase';
import AlertMessage from '../components/AlertMessage';
import { LogInRequest } from '../gen/pb/auth_pb';
import { AuthClient } from '../gen/pb/auth_grpc_web_pb';
import { useHistory } from "react-router-dom"

const useStyles = makeStyles((theme) => ({
    paper: {
        marginTop: theme.spacing(8),
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
    },
    avatar: {
        margin: theme.spacing(1),
        backgroundColor: theme.palette.secondary.main,
    },
    form: {
        width: '100%',
        marginTop: theme.spacing(3),
    },
    submit: {
        margin: theme.spacing(3, 0, 2),
    },
}));

const LogIn = () => {
    const classes = useStyles();
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [errors, setErrors] = useState({});
    const [authError, setAuthError] = useState("");
    const history = useHistory();

    /*
     * ログインボタン押下時のイベントハンドラ
     */
    const handleSubmit = e => {
        e.preventDefault();
        setErrors({})

        const validationErrors = validate();

        if (Object.keys(validationErrors).length > 0) {
            setErrors(validationErrors);
            return;
        }
        login();
    }

    /*
     * ログイン
     */
    const login = () => {
        const request = new LogInRequest()
        request.setEmail(email)
        request.setPassword(password)

        const client = new AuthClient("http://localhost:8080")
        client.logIn(request, {}, (err, response) => {
            if (err) {
                setAuthError(err.message);
                return;
            }
            const token = response.getToken();
            auth.signInWithCustomToken(token).catch(err => {
                setAuthError(err.message);
                return;
            });
            history.push("/");
        });
    }

    /*
     * 入力バリデーション
     */
    const validate = () => {
        const errors = {};
        if (!email) {
            errors.email = "メールアドレスを入力してください";
        }
        if (!password) {
            errors.password = "パスワードを入力してください";
        }
        return errors;
    }

    return (
        <Container component="main" maxWidth="xs">
            <CssBaseline />
            <div className={classes.paper}>
                <Avatar className={classes.avatar}>
                    <LockOutlinedIcon />
                </Avatar>
                <Typography component="h1" variant="h5">
                    ログイン
                </Typography>
                <form className={classes.form} noValidate onSubmit={e => handleSubmit(e)}>
                    <Grid item xs={12}>
                        <TextField
                            variant="outlined"
                            required
                            fullWidth
                            id="email"
                            label="Email Address"
                            name="email"
                            autoComplete="email"
                            onChange={e => setEmail(e.target.value)}
                            value={email}
                        />
                        {errors.email && <div>{errors.email}</div>}
                    </Grid>
                    <Grid item xs={12}>
                        <TextField
                            variant="outlined"
                            required
                            fullWidth
                            name="password"
                            label="Password"
                            type="password"
                            id="password"
                            autoComplete="current-password"
                            onChange={e => setPassword(e.target.value)}
                        />
                        {errors.password && <div>{errors.password}</div>}
                    </Grid>
                    {authError && <AlertMessage message={authError} />}
                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        color="primary"
                        className={classes.submit}
                    >
                        ログイン
                    </Button>
                    <Grid container justify="flex-end">
                        <Grid item>
                            <Link href="/signup" variant="body2">
                                アカウントをお持ちでない方はこちら
                            </Link>
                        </Grid>
                    </Grid>
                </form>
            </div>
        </Container >
    );
}

export default LogIn;